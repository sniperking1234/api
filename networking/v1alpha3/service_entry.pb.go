// Copyright 2018 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: networking/v1alpha3/service_entry.proto

// $schema: istio.networking.v1alpha3.ServiceEntry
// $title: Service Entry
// $description: Configuration affecting service registry.
// $location: https://istio.io/docs/reference/config/networking/service-entry.html
// $aliases: [/docs/reference/config/networking/v1alpha3/service-entry]

// `ServiceEntry` enables adding additional entries into Istio's
// internal service registry, so that auto-discovered services in the
// mesh can access/route to these manually specified services. A
// service entry describes the properties of a service (DNS name,
// VIPs, ports, protocols, endpoints). These services could be
// external to the mesh (e.g., web APIs) or mesh-internal services
// that are not part of the platform's service registry (e.g., a set
// of VMs talking to services in Kubernetes). In addition, the
// endpoints of a service entry can also be dynamically selected by
// using the `workloadSelector` field. These endpoints can be VM
// workloads declared using the `WorkloadEntry` object or Kubernetes
// pods. The ability to select both pods and VMs under a single
// service allows for migration of services from VMs to Kubernetes
// without having to change the existing DNS names associated with the
// services.
//
// The following example declares a few external APIs accessed by internal
// applications over HTTPS. The sidecar inspects the SNI value in the
// ClientHello message to route to the appropriate external service.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-https
// spec:
//   hosts:
//   - api.dropboxapi.com
//   - www.googleapis.com
//   - api.facebook.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: DNS
// ```
//
// The following configuration adds a set of MongoDB instances running on
// unmanaged VMs to Istio's registry, so that these services can be treated
// as any other service in the mesh. The associated DestinationRule is used
// to initiate mTLS connections to the database instances.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-mongocluster
// spec:
//   hosts:
//   - mymongodb.somedomain # not used
//   addresses:
//   - 192.192.192.192/24 # VIPs
//   ports:
//   - number: 27018
//     name: mongodb
//     protocol: MONGO
//   location: MESH_INTERNAL
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
// ```
//
// and the associated DestinationRule
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: mtls-mongocluster
// spec:
//   host: mymongodb.somedomain
//   trafficPolicy:
//     tls:
//       mode: MUTUAL
//       clientCertificate: /etc/certs/myclientcert.pem
//       privateKey: /etc/certs/client_private_key.pem
//       caCertificates: /etc/certs/rootcacerts.pem
// ```
//
// The following example uses a combination of service entry and TLS
// routing in a virtual service to steer traffic based on the SNI value to
// an internal egress firewall.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-redirect
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: NONE
// ```
//
// And the associated VirtualService to route based on the SNI value.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: tls-routing
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   tls:
//   - match:
//     - sniHosts:
//       - wikipedia.org
//       - "*.wikipedia.org"
//     route:
//     - destination:
//         host: internal-egress-firewall.ns1.svc.cluster.local
// ```
//
// The virtual service with TLS match serves to override the default SNI
// match. In the absence of a virtual service, traffic will be forwarded to
// the wikipedia domains.
//
// The following example demonstrates the use of a dedicated egress gateway
// through which all external service traffic is forwarded.
// The 'exportTo' field allows for control over the visibility of a service
// declaration to other namespaces in the mesh. By default, a service is exported
// to all namespaces. The following example restricts the visibility to the
// current namespace, represented by ".", so that it cannot be used by other
// namespaces.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-httpbin
//   namespace : egress
// spec:
//   hosts:
//   - example.com
//   exportTo:
//   - "."
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
// ```
//
// Define a gateway to handle all egress traffic.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//  name: istio-egressgateway
//  namespace: istio-system
// spec:
//  selector:
//    istio: egressgateway
//  servers:
//  - port:
//      number: 80
//      name: http
//      protocol: HTTP
//    hosts:
//    - "*"
// ```
//
// And the associated `VirtualService` to route from the sidecar to the
// gateway service (`istio-egressgateway.istio-system.svc.cluster.local`), as
// well as route from the gateway to the external service. Note that the
// virtual service is exported to all namespaces enabling them to route traffic
// through the gateway to the external service. Forcing traffic to go through
// a managed middle proxy like this is a common practice.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: gateway-routing
//   namespace: egress
// spec:
//   hosts:
//   - example.com
//   exportTo:
//   - "*"
//   gateways:
//   - mesh
//   - istio-egressgateway
//   http:
//   - match:
//     - port: 80
//       gateways:
//       - mesh
//     route:
//     - destination:
//         host: istio-egressgateway.istio-system.svc.cluster.local
//   - match:
//     - port: 80
//       gateways:
//       - istio-egressgateway
//     route:
//     - destination:
//         host: example.com
// ```
//
// The following example demonstrates the use of wildcards in the hosts for
// external services. If the connection has to be routed to the IP address
// requested by the application (i.e. application resolves DNS and attempts
// to connect to a specific IP), the resolution mode must be set to `NONE`.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-wildcard-example
// spec:
//   hosts:
//   - "*.bar.com"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: NONE
// ```
//
// The following example demonstrates a service that is available via a
// Unix Domain Socket on the host of the client. The resolution must be
// set to STATIC to use Unix address endpoints.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: unix-domain-socket-example
// spec:
//   hosts:
//   - "example.unix.local"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: unix:///var/run/example/socket
// ```
//
// For HTTP-based services, it is possible to create a `VirtualService`
// backed by multiple DNS addressable endpoints. In such a scenario, the
// application can use the `HTTP_PROXY` environment variable to transparently
// reroute API calls for the `VirtualService` to a chosen backend. For
// example, the following configuration creates a non-existent external
// service called foo.bar.com backed by three domains: us.foo.bar.com:8080,
// uk.foo.bar.com:9080, and in.foo.bar.com:7080
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-dns
// spec:
//   hosts:
//   - foo.bar.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
//   endpoints:
//   - address: us.foo.bar.com
//     ports:
//       http: 8080
//   - address: uk.foo.bar.com
//     ports:
//       http: 9080
//   - address: in.foo.bar.com
//     ports:
//       http: 7080
// ```
//
// With `HTTP_PROXY=http://localhost/`, calls from the application to
// `http://foo.bar.com` will be load balanced across the three domains
// specified above. In other words, a call to `http://foo.bar.com/baz` would
// be translated to `http://uk.foo.bar.com/baz`.
//
// The following example illustrates the usage of a `ServiceEntry`
// containing a subject alternate name
// whose format conforms to the [SPIFFE standard](https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md):
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: httpbin
//   namespace : httpbin-ns
// spec:
//   hosts:
//   - example.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
//   subjectAltNames:
//   - "spiffe://cluster.local/ns/httpbin-ns/sa/httpbin-service-account"
// ```
//
// The following example demonstrates the use of `ServiceEntry` with a
// `workloadSelector` to handle the migration of a service
// `details.bookinfo.com` from VMs to Kubernetes. The service has two
// VM-based instances with sidecars as well as a set of Kubernetes
// pods managed by a standard deployment object. Consumers of this
// service in the mesh will be automatically load balanced across the
// VMs and Kubernetes.
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-vm-1
// spec:
//   serviceAccount: details
//   address: 2.2.2.2
//   labels:
//     app: details
//     instance-id: vm1
// ---
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-vm-2
// spec:
//   serviceAccount: details
//   address: 3.3.3.3
//   labels:
//     app: details
//     instance-id: vm2
// ```
//
// Assuming there is also a Kubernetes deployment with pod labels
// `app: details` using the same service account `details`, the
// following service entry declares a service spanning both VMs and
// Kubernetes:
//
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details
// ```

package v1alpha3

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Location specifies whether the service is part of Istio mesh or
// outside the mesh.  Location determines the behavior of several
// features, such as service-to-service mTLS authentication, policy
// enforcement, etc. When communicating with services outside the mesh,
// Istio's mTLS authentication is disabled, and policy enforcement is
// performed on the client-side as opposed to server-side.
type ServiceEntry_Location int32

const (
	// Signifies that the service is external to the mesh. Typically used
	// to indicate external services consumed through APIs.
	ServiceEntry_MESH_EXTERNAL ServiceEntry_Location = 0
	// Signifies that the service is part of the mesh. Typically used to
	// indicate services added explicitly as part of expanding the service
	// mesh to include unmanaged infrastructure (e.g., VMs added to a
	// Kubernetes based service mesh).
	ServiceEntry_MESH_INTERNAL ServiceEntry_Location = 1
)

// Enum value maps for ServiceEntry_Location.
var (
	ServiceEntry_Location_name = map[int32]string{
		0: "MESH_EXTERNAL",
		1: "MESH_INTERNAL",
	}
	ServiceEntry_Location_value = map[string]int32{
		"MESH_EXTERNAL": 0,
		"MESH_INTERNAL": 1,
	}
)

func (x ServiceEntry_Location) Enum() *ServiceEntry_Location {
	p := new(ServiceEntry_Location)
	*p = x
	return p
}

func (x ServiceEntry_Location) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceEntry_Location) Descriptor() protoreflect.EnumDescriptor {
	return file_networking_v1alpha3_service_entry_proto_enumTypes[0].Descriptor()
}

func (ServiceEntry_Location) Type() protoreflect.EnumType {
	return &file_networking_v1alpha3_service_entry_proto_enumTypes[0]
}

func (x ServiceEntry_Location) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceEntry_Location.Descriptor instead.
func (ServiceEntry_Location) EnumDescriptor() ([]byte, []int) {
	return file_networking_v1alpha3_service_entry_proto_rawDescGZIP(), []int{0, 0}
}

// Resolution determines how the proxy will resolve the IP addresses of
// the network endpoints associated with the service, so that it can
// route to one of them. The resolution mode specified here has no impact
// on how the application resolves the IP address associated with the
// service. The application may still have to use DNS to resolve the
// service to an IP so that the outbound traffic can be captured by the
// Proxy. Alternatively, for HTTP services, the application could
// directly communicate with the proxy (e.g., by setting HTTP_PROXY) to
// talk to these services.
type ServiceEntry_Resolution int32

const (
	// Assume that incoming connections have already been resolved (to a
	// specific destination IP address). Such connections are typically
	// routed via the proxy using mechanisms such as IP table REDIRECT/
	// eBPF. After performing any routing related transformations, the
	// proxy will forward the connection to the IP address to which the
	// connection was bound.
	ServiceEntry_NONE ServiceEntry_Resolution = 0
	// Use the static IP addresses specified in endpoints (see below) as the
	// backing instances associated with the service.
	ServiceEntry_STATIC ServiceEntry_Resolution = 1
	// Attempt to resolve the IP address by querying the ambient DNS,
	// asynchronously. If no endpoints are specified, the proxy
	// will resolve the DNS address specified in the hosts field, if
	// wildcards are not used. If endpoints are specified, the DNS
	// addresses specified in the endpoints will be resolved to determine
	// the destination IP address.  DNS resolution cannot be used with Unix
	// domain socket endpoints.
	ServiceEntry_DNS ServiceEntry_Resolution = 2
	// Attempt to resolve the IP address by querying the ambient DNS,
	// asynchronously. Unlike `DNS`, `DNS_ROUND_ROBIN` only uses the
	// first IP address returned when a new connection needs to be initiated
	// without relying on complete results of DNS resolution, and connections
	// made to hosts will be retained even if DNS records change frequently
	// eliminating draining connection pools and connection cycling.
	// This is best suited for large web scale services that
	// must be accessed via DNS. The proxy will resolve the DNS address
	// specified in the hosts field, if wildcards are not used. DNS resolution
	// cannot be used with Unix domain socket endpoints.
	ServiceEntry_DNS_ROUND_ROBIN ServiceEntry_Resolution = 3
)

// Enum value maps for ServiceEntry_Resolution.
var (
	ServiceEntry_Resolution_name = map[int32]string{
		0: "NONE",
		1: "STATIC",
		2: "DNS",
		3: "DNS_ROUND_ROBIN",
	}
	ServiceEntry_Resolution_value = map[string]int32{
		"NONE":            0,
		"STATIC":          1,
		"DNS":             2,
		"DNS_ROUND_ROBIN": 3,
	}
)

func (x ServiceEntry_Resolution) Enum() *ServiceEntry_Resolution {
	p := new(ServiceEntry_Resolution)
	*p = x
	return p
}

func (x ServiceEntry_Resolution) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceEntry_Resolution) Descriptor() protoreflect.EnumDescriptor {
	return file_networking_v1alpha3_service_entry_proto_enumTypes[1].Descriptor()
}

func (ServiceEntry_Resolution) Type() protoreflect.EnumType {
	return &file_networking_v1alpha3_service_entry_proto_enumTypes[1]
}

func (x ServiceEntry_Resolution) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceEntry_Resolution.Descriptor instead.
func (ServiceEntry_Resolution) EnumDescriptor() ([]byte, []int) {
	return file_networking_v1alpha3_service_entry_proto_rawDescGZIP(), []int{0, 1}
}

// ServiceEntry enables adding additional entries into Istio's internal
// service registry.
//
// <!-- crd generation tags
// +cue-gen:ServiceEntry:groupName:networking.istio.io
// +cue-gen:ServiceEntry:versions:v1beta1,v1alpha3,v1
// +cue-gen:ServiceEntry:annotations:helm.sh/resource-policy=keep
// +cue-gen:ServiceEntry:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:ServiceEntry:subresource:status
// +cue-gen:ServiceEntry:scope:Namespaced
// +cue-gen:ServiceEntry:resource:categories=istio-io,networking-istio-io,shortNames=se,plural=serviceentries
// +cue-gen:ServiceEntry:printerColumn:name=Hosts,type=string,JSONPath=.spec.hosts,description="The hosts associated with the ServiceEntry"
// +cue-gen:ServiceEntry:printerColumn:name=Location,type=string,JSONPath=.spec.location,description="Whether the service is external to the
// mesh or part of the mesh (MESH_EXTERNAL or MESH_INTERNAL)"
// +cue-gen:ServiceEntry:printerColumn:name=Resolution,type=string,JSONPath=.spec.resolution,description="Service resolution mode for the hosts
// (NONE, STATIC, or DNS)"
// +cue-gen:ServiceEntry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:ServiceEntry:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
// <!-- istio code generation tags
// +istio.io/sync-start
// -->
type ServiceEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hosts associated with the ServiceEntry. Could be a DNS
	// name with wildcard prefix.
	//
	// 1. The hosts field is used to select matching hosts in VirtualServices and DestinationRules.
	// 2. For HTTP traffic the HTTP Host/Authority header will be matched against the hosts field.
	// 3. For HTTPs or TLS traffic containing Server Name Indication (SNI), the SNI value
	// will be matched against the hosts field.
	//
	// **NOTE 1:** When resolution is set to type DNS and no endpoints
	// are specified, the host field will be used as the DNS name of the
	// endpoint to route traffic to.
	//
	// **NOTE 2:** If the hostname matches with the name of a service
	// from another service registry such as Kubernetes that also
	// supplies its own set of endpoints, the ServiceEntry will be
	// treated as a decorator of the existing Kubernetes
	// service. Properties in the service entry will be added to the
	// Kubernetes service if applicable. Currently, only the following
	// additional properties will be considered by `istiod`:
	//
	//  1. subjectAltNames: In addition to verifying the SANs of the
	//     service accounts associated with the pods of the service, the
	//     SANs specified here will also be verified.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// The virtual IP addresses associated with the service. Could be CIDR
	// prefix. For HTTP traffic, generated route configurations will include http route
	// domains for both the `addresses` and `hosts` field values and the destination will
	// be identified based on the HTTP Host/Authority header.
	// If one or more IP addresses are specified,
	// the incoming traffic will be identified as belonging to this service
	// if the destination IP matches the IP/CIDRs specified in the addresses
	// field. If the Addresses field is empty, traffic will be identified
	// solely based on the destination port. In such scenarios, the port on
	// which the service is being accessed must not be shared by any other
	// service in the mesh. In other words, the sidecar will behave as a
	// simple TCP proxy, forwarding incoming traffic on a specified port to
	// the specified destination endpoint IP/host. Unix domain socket
	// addresses are not supported in this field.
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// The ports associated with the external service. If the
	// Endpoints are Unix domain socket addresses, there must be exactly one
	// port.
	Ports []*ServicePort `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	// Specify whether the service should be considered external to the mesh
	// or part of the mesh.
	Location ServiceEntry_Location `protobuf:"varint,4,opt,name=location,proto3,enum=istio.networking.v1alpha3.ServiceEntry_Location" json:"location,omitempty"`
	// Service resolution mode for the hosts. Care must be taken
	// when setting the resolution mode to NONE for a TCP port without
	// accompanying IP addresses. In such cases, traffic to any IP on
	// said port will be allowed (i.e. `0.0.0.0:<port>`).
	Resolution ServiceEntry_Resolution `protobuf:"varint,5,opt,name=resolution,proto3,enum=istio.networking.v1alpha3.ServiceEntry_Resolution" json:"resolution,omitempty"`
	// One or more endpoints associated with the service. Only one of
	// `endpoints` or `workloadSelector` can be specified.
	Endpoints []*WorkloadEntry `protobuf:"bytes,6,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Applicable only for MESH_INTERNAL services. Only one of
	// `endpoints` or `workloadSelector` can be specified. Selects one
	// or more Kubernetes pods or VM workloads (specified using
	// `WorkloadEntry`) based on their labels. The `WorkloadEntry` object
	// representing the VMs should be defined in the same namespace as
	// the ServiceEntry.
	WorkloadSelector *WorkloadSelector `protobuf:"bytes,9,opt,name=workload_selector,json=workloadSelector,proto3" json:"workload_selector,omitempty"`
	// A list of namespaces to which this service is exported. Exporting a service
	// allows it to be used by sidecars, gateways and virtual services defined in
	// other namespaces. This feature provides a mechanism for service owners
	// and mesh administrators to control the visibility of services across
	// namespace boundaries.
	//
	// If no namespaces are specified then the service is exported to all
	// namespaces by default.
	//
	// The value "." is reserved and defines an export to the same namespace that
	// the service is declared in. Similarly the value "*" is reserved and
	// defines an export to all namespaces.
	//
	// For a Kubernetes Service, the equivalent effect can be achieved by setting
	// the annotation "networking.istio.io/exportTo" to a comma-separated list
	// of namespace names.
	ExportTo []string `protobuf:"bytes,7,rep,name=export_to,json=exportTo,proto3" json:"export_to,omitempty"`
	// If specified, the proxy will verify that the server certificate's
	// subject alternate name matches one of the specified values.
	//
	// NOTE: When using the workloadEntry with workloadSelectors, the
	// service account specified in the workloadEntry will also be used
	// to derive the additional subject alternate names that should be
	// verified.
	SubjectAltNames []string `protobuf:"bytes,8,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
}

func (x *ServiceEntry) Reset() {
	*x = ServiceEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networking_v1alpha3_service_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceEntry) ProtoMessage() {}

func (x *ServiceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_networking_v1alpha3_service_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceEntry.ProtoReflect.Descriptor instead.
func (*ServiceEntry) Descriptor() ([]byte, []int) {
	return file_networking_v1alpha3_service_entry_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceEntry) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *ServiceEntry) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *ServiceEntry) GetPorts() []*ServicePort {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *ServiceEntry) GetLocation() ServiceEntry_Location {
	if x != nil {
		return x.Location
	}
	return ServiceEntry_MESH_EXTERNAL
}

func (x *ServiceEntry) GetResolution() ServiceEntry_Resolution {
	if x != nil {
		return x.Resolution
	}
	return ServiceEntry_NONE
}

func (x *ServiceEntry) GetEndpoints() []*WorkloadEntry {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ServiceEntry) GetWorkloadSelector() *WorkloadSelector {
	if x != nil {
		return x.WorkloadSelector
	}
	return nil
}

func (x *ServiceEntry) GetExportTo() []string {
	if x != nil {
		return x.ExportTo
	}
	return nil
}

func (x *ServiceEntry) GetSubjectAltNames() []string {
	if x != nil {
		return x.SubjectAltNames
	}
	return nil
}

// ServicePort describes the properties of a specific port of a service.
type ServicePort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A valid non-negative integer port number.
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// The protocol exposed on the port.
	// MUST BE one of HTTP|HTTPS|GRPC|HTTP2|MONGO|TCP|TLS.
	// TLS implies the connection will be routed based on the SNI header to
	// the destination without terminating the TLS connection.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Label assigned to the port.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The port number on the endpoint where the traffic will be
	// received. If unset, default to `number`.
	TargetPort uint32 `protobuf:"varint,4,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
}

func (x *ServicePort) Reset() {
	*x = ServicePort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networking_v1alpha3_service_entry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePort) ProtoMessage() {}

func (x *ServicePort) ProtoReflect() protoreflect.Message {
	mi := &file_networking_v1alpha3_service_entry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePort.ProtoReflect.Descriptor instead.
func (*ServicePort) Descriptor() ([]byte, []int) {
	return file_networking_v1alpha3_service_entry_proto_rawDescGZIP(), []int{1}
}

func (x *ServicePort) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ServicePort) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServicePort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServicePort) GetTargetPort() uint32 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

var File_networking_v1alpha3_service_entry_proto protoreflect.FileDescriptor

var file_networking_v1alpha3_service_entry_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x05, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x58, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x33, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x10, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6c, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41,
	0x6c, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45, 0x53, 0x48, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45, 0x53, 0x48, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x22, 0x40, 0x0a, 0x0a, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x44, 0x4e, 0x53, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x4e, 0x53, 0x5f, 0x52, 0x4f,
	0x55, 0x4e, 0x44, 0x5f, 0x52, 0x4f, 0x42, 0x49, 0x4e, 0x10, 0x03, 0x22, 0x80, 0x01, 0x0a, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x22,
	0x5a, 0x20, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_networking_v1alpha3_service_entry_proto_rawDescOnce sync.Once
	file_networking_v1alpha3_service_entry_proto_rawDescData = file_networking_v1alpha3_service_entry_proto_rawDesc
)

func file_networking_v1alpha3_service_entry_proto_rawDescGZIP() []byte {
	file_networking_v1alpha3_service_entry_proto_rawDescOnce.Do(func() {
		file_networking_v1alpha3_service_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_networking_v1alpha3_service_entry_proto_rawDescData)
	})
	return file_networking_v1alpha3_service_entry_proto_rawDescData
}

var file_networking_v1alpha3_service_entry_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_networking_v1alpha3_service_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_networking_v1alpha3_service_entry_proto_goTypes = []interface{}{
	(ServiceEntry_Location)(0),   // 0: istio.networking.v1alpha3.ServiceEntry.Location
	(ServiceEntry_Resolution)(0), // 1: istio.networking.v1alpha3.ServiceEntry.Resolution
	(*ServiceEntry)(nil),         // 2: istio.networking.v1alpha3.ServiceEntry
	(*ServicePort)(nil),          // 3: istio.networking.v1alpha3.ServicePort
	(*WorkloadEntry)(nil),        // 4: istio.networking.v1alpha3.WorkloadEntry
	(*WorkloadSelector)(nil),     // 5: istio.networking.v1alpha3.WorkloadSelector
}
var file_networking_v1alpha3_service_entry_proto_depIdxs = []int32{
	3, // 0: istio.networking.v1alpha3.ServiceEntry.ports:type_name -> istio.networking.v1alpha3.ServicePort
	0, // 1: istio.networking.v1alpha3.ServiceEntry.location:type_name -> istio.networking.v1alpha3.ServiceEntry.Location
	1, // 2: istio.networking.v1alpha3.ServiceEntry.resolution:type_name -> istio.networking.v1alpha3.ServiceEntry.Resolution
	4, // 3: istio.networking.v1alpha3.ServiceEntry.endpoints:type_name -> istio.networking.v1alpha3.WorkloadEntry
	5, // 4: istio.networking.v1alpha3.ServiceEntry.workload_selector:type_name -> istio.networking.v1alpha3.WorkloadSelector
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_networking_v1alpha3_service_entry_proto_init() }
func file_networking_v1alpha3_service_entry_proto_init() {
	if File_networking_v1alpha3_service_entry_proto != nil {
		return
	}
	file_networking_v1alpha3_sidecar_proto_init()
	file_networking_v1alpha3_workload_entry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_networking_v1alpha3_service_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_networking_v1alpha3_service_entry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_networking_v1alpha3_service_entry_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_networking_v1alpha3_service_entry_proto_goTypes,
		DependencyIndexes: file_networking_v1alpha3_service_entry_proto_depIdxs,
		EnumInfos:         file_networking_v1alpha3_service_entry_proto_enumTypes,
		MessageInfos:      file_networking_v1alpha3_service_entry_proto_msgTypes,
	}.Build()
	File_networking_v1alpha3_service_entry_proto = out.File
	file_networking_v1alpha3_service_entry_proto_rawDesc = nil
	file_networking_v1alpha3_service_entry_proto_goTypes = nil
	file_networking_v1alpha3_service_entry_proto_depIdxs = nil
}
