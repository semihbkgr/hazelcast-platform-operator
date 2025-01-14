package hazelcast

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/go-logr/logr"
	proto "github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/cluster"
	hztypes "github.com/hazelcast/hazelcast-go-client/types"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/event"

	hazelcastv1alpha1 "github.com/hazelcast/hazelcast-platform-operator/api/v1alpha1"
	hzclient "github.com/hazelcast/hazelcast-platform-operator/internal/hazelcast-client"
	n "github.com/hazelcast/hazelcast-platform-operator/internal/naming"
	codecTypes "github.com/hazelcast/hazelcast-platform-operator/internal/protocol/types"
)

func fakeK8sClient(initObjs ...client.Object) client.Client {
	scheme, _ := hazelcastv1alpha1.SchemeBuilder.
		Register(&hazelcastv1alpha1.Hazelcast{}, &hazelcastv1alpha1.HazelcastList{}, &v1.ClusterRole{}, &v1.ClusterRoleBinding{},
			&hazelcastv1alpha1.Cache{}, &hazelcastv1alpha1.CacheList{}, &corev1.ConfigMap{}).
		Build()
	return fake.NewClientBuilder().WithScheme(scheme).WithObjects(initObjs...).Build()
}

func fakeHttpServer(url string, handler http.HandlerFunc) (*httptest.Server, error) {
	l, err := net.Listen("tcp", url)
	if err != nil {
		return nil, err
	}
	ts := httptest.NewUnstartedServer(handler)
	_ = ts.Listener.Close()
	ts.Listener = l
	ts.Start()
	return ts, nil
}

type fakeHzClientRegistry struct {
	Clients sync.Map
}

func (cr *fakeHzClientRegistry) GetOrCreate(ctx context.Context, nn types.NamespacedName) (hzclient.Client, error) {
	client, ok := cr.Get(nn)
	if !ok {
		return client, fmt.Errorf("Fake client was not set before test")
	}
	return client, nil
}

func (cr *fakeHzClientRegistry) Set(ns types.NamespacedName, cl hzclient.Client) {
	cr.Clients.Store(ns, cl)
}

func (cr *fakeHzClientRegistry) Get(ns types.NamespacedName) (hzclient.Client, bool) {
	if v, ok := cr.Clients.Load(ns); ok {
		return v.(hzclient.Client), true
	}
	return nil, false
}

func (cr *fakeHzClientRegistry) Delete(ctx context.Context, ns types.NamespacedName) error {
	if c, ok := cr.Clients.LoadAndDelete(ns); ok {
		return c.(hzclient.Client).Shutdown(ctx) //nolint:errcheck
	}
	return nil
}

type fakeHttpClientRegistry struct {
	clients sync.Map
}

func (hr *fakeHttpClientRegistry) Create(_ context.Context, _ client.Client, ns string) (*http.Client, error) {
	if v, ok := hr.clients.Load(types.NamespacedName{Name: n.MTLSCertSecretName, Namespace: ns}); ok {
		return v.(*http.Client), nil
	}
	return nil, errors.New("no client found")
}

func (hr *fakeHttpClientRegistry) Get(ns string) (*http.Client, bool) {
	if v, ok := hr.clients.Load(types.NamespacedName{Name: n.MTLSCertSecretName, Namespace: ns}); ok {
		return v.(*http.Client), ok
	}
	return nil, false
}

func (hr *fakeHttpClientRegistry) Delete(ns string) {
	hr.clients.Delete(types.NamespacedName{Name: n.MTLSCertSecretName, Namespace: ns})
}

func (hr *fakeHttpClientRegistry) Set(ns string, cl *http.Client) {
	hr.clients.Store(types.NamespacedName{Name: n.MTLSCertSecretName, Namespace: ns}, cl)
}

type fakeHzClient struct {
	tOrderedMembers          []cluster.MemberInfo
	tIsClientConnected       bool
	tAreAllMembersAccessible bool
	tRunning                 bool
	tInvokeOnMember          func(ctx context.Context, req *proto.ClientMessage, uuid hztypes.UUID, opts *proto.InvokeOptions) (*proto.ClientMessage, error)
	tInvokeOnRandomTarget    func(ctx context.Context, req *proto.ClientMessage, opts *proto.InvokeOptions) (*proto.ClientMessage, error)
	tShutDown                error
	tUUID                    hztypes.UUID
}

func (cl *fakeHzClient) OrderedMembers() []cluster.MemberInfo {
	return cl.tOrderedMembers
}

func (cl *fakeHzClient) IsClientConnected() bool {
	return cl.tIsClientConnected
}

func (cl *fakeHzClient) AreAllMembersAccessible() bool {
	return cl.tAreAllMembersAccessible
}

func (cl *fakeHzClient) InvokeOnMember(ctx context.Context, req *proto.ClientMessage, uuid hztypes.UUID, opts *proto.InvokeOptions) (*proto.ClientMessage, error) {
	if cl.tInvokeOnMember != nil {
		return cl.tInvokeOnMember(ctx, req, uuid, opts)
	}
	return nil, nil
}

func (cl *fakeHzClient) InvokeOnRandomTarget(ctx context.Context, req *proto.ClientMessage, opts *proto.InvokeOptions) (*proto.ClientMessage, error) {
	if cl.tInvokeOnRandomTarget != nil {
		return cl.tInvokeOnRandomTarget(ctx, req, opts)
	}
	return nil, nil
}

func (cl *fakeHzClient) Running() bool {
	return cl.tRunning
}

func (cl *fakeHzClient) ClusterId() hztypes.UUID {
	return cl.tUUID
}

func (cl *fakeHzClient) Shutdown(_ context.Context) error {
	return cl.tShutDown
}

type fakeHzStatusServiceRegistry struct {
	statusServices sync.Map
}

func (ssr *fakeHzStatusServiceRegistry) Create(ns types.NamespacedName, _ hzclient.Client, l logr.Logger, channel chan event.GenericEvent) hzclient.StatusService {
	ss, ok := ssr.Get(ns)
	if ok {
		return ss
	}
	ssr.statusServices.Store(ns, ss)
	ss.Start()
	return ss
}

func (ssr *fakeHzStatusServiceRegistry) Set(ns types.NamespacedName, ss hzclient.StatusService) {
	ssr.statusServices.Store(ns, ss)
}

func (ssr *fakeHzStatusServiceRegistry) Get(ns types.NamespacedName) (hzclient.StatusService, bool) {
	if v, ok := ssr.statusServices.Load(ns); ok {
		return v.(hzclient.StatusService), ok
	}
	return nil, false
}

func (ssr *fakeHzStatusServiceRegistry) Delete(ns types.NamespacedName) {
	if ss, ok := ssr.statusServices.LoadAndDelete(ns); ok {
		ss.(hzclient.StatusService).Stop()
	}
}

type fakeHzStatusService struct {
	Status              *hzclient.Status
	timedMemberStateMap map[hztypes.UUID]*codecTypes.TimedMemberStateWrapper
	tStart              func()
	tUpdateMembers      func(ss *fakeHzStatusService, ctx context.Context)
}

func (ss *fakeHzStatusService) Start() {
	if ss.tStart != nil {
		ss.tStart()
	}
}

func (ss *fakeHzStatusService) GetStatus() *hzclient.Status {
	return ss.Status
}

func (ss *fakeHzStatusService) UpdateMembers(ctx context.Context) {
	if ss.tUpdateMembers != nil {
		ss.tUpdateMembers(ss, ctx)
	}
}

func (ss *fakeHzStatusService) GetTimedMemberState(_ context.Context, uuid hztypes.UUID) (*codecTypes.TimedMemberStateWrapper, error) {
	if state, ok := ss.timedMemberStateMap[uuid]; ok {
		return state, nil
	}
	return nil, nil
}

func (ss *fakeHzStatusService) Stop() {
}
