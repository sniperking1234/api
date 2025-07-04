// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1beta1

import "istio.io/api/networking/v1alpha3"

// DestinationRule defines policies that apply to traffic intended for a service
// after routing has occurred.
//
// <!-- crd generation tags
// +cue-gen:DestinationRule:groupName:networking.istio.io
// +cue-gen:DestinationRule:versions:v1,v1beta1,v1alpha3
// +cue-gen:DestinationRule:annotations:helm.sh/resource-policy=keep
// +cue-gen:DestinationRule:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:DestinationRule:subresource:status
// +cue-gen:DestinationRule:scope:Namespaced
// +cue-gen:DestinationRule:resource:categories=istio-io,networking-istio-io,shortNames=dr
// +cue-gen:DestinationRule:printerColumn:name=Host,type=string,JSONPath=.spec.host,description="The name of a service from the service registry"
// +cue-gen:DestinationRule:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. For more information, see [Kubernetes API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)"
// +cue-gen:DestinationRule:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
type DestinationRule = v1alpha3.DestinationRule

// Traffic policies to apply for a specific destination, across all
// destination ports. See DestinationRule for examples.
type TrafficPolicy = v1alpha3.TrafficPolicy

// Traffic policies that apply to specific ports of the service
type TrafficPolicy_PortTrafficPolicy = v1alpha3.TrafficPolicy_PortTrafficPolicy
type TrafficPolicy_TunnelSettings = v1alpha3.TrafficPolicy_TunnelSettings
type TrafficPolicy_ProxyProtocol = v1alpha3.TrafficPolicy_ProxyProtocol
type TrafficPolicy_ProxyProtocol_VERSION = v1alpha3.TrafficPolicy_ProxyProtocol_VERSION

// ⁣PROXY protocol version 1. Human readable format.
const TrafficPolicy_ProxyProtocol_V1 TrafficPolicy_ProxyProtocol_VERSION = v1alpha3.TrafficPolicy_ProxyProtocol_V1

// ⁣PROXY protocol version 2. Binary format.
const TrafficPolicy_ProxyProtocol_V2 TrafficPolicy_ProxyProtocol_VERSION = v1alpha3.TrafficPolicy_ProxyProtocol_V2

type TrafficPolicy_RetryBudget = v1alpha3.TrafficPolicy_RetryBudget

// A subset of endpoints of a service. Subsets can be used for scenarios
// like A/B testing, or routing to a specific version of a service. Refer
// to [VirtualService](https://istio.io/docs/reference/config/networking/virtual-service/#VirtualService) documentation for examples of using
// subsets in these scenarios. In addition, traffic policies defined at the
// service-level can be overridden at a subset-level. The following rule
// uses a round robin load balancing policy for all traffic going to a
// subset named testversion that is composed of endpoints (e.g., pods) with
// labels (version:v3).
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: bookinfo-ratings
//
// spec:
//
//	host: ratings.prod.svc.cluster.local
//	trafficPolicy:
//	  loadBalancer:
//	    simple: LEAST_REQUEST
//	subsets:
//	- name: testversion
//	  labels:
//	    version: v3
//	  trafficPolicy:
//	    loadBalancer:
//	      simple: ROUND_ROBIN
//
// ```
//
// **Note:** Policies specified for subsets will not take effect until
// a route rule explicitly sends traffic to this subset.
//
// One or more labels are typically required to identify the subset destination,
// however, when the corresponding DestinationRule represents a host that
// supports multiple SNI hosts (e.g., an egress gateway), a subset without labels
// may be meaningful. In this case a traffic policy with [ClientTLSSettings](#ClientTLSSettings)
// can be used to identify a specific SNI host corresponding to the named subset.
type Subset = v1alpha3.Subset

// Load balancing policies to apply for a specific destination. See Envoy's
// load balancing
// [documentation](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/load_balancing)
// for more details.
//
// For example, the following rule uses a round robin load balancing policy
// for all traffic going to the ratings service.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: bookinfo-ratings
//
// spec:
//
//	host: ratings.prod.svc.cluster.local
//	trafficPolicy:
//	  loadBalancer:
//	    simple: ROUND_ROBIN
//
// ```
//
// The following example sets up sticky sessions for the ratings service
// hashing-based load balancer for the same ratings service using the
// the User cookie as the hash key.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: bookinfo-ratings
//
// spec:
//
//	host: ratings.prod.svc.cluster.local
//	trafficPolicy:
//	  loadBalancer:
//	    consistentHash:
//	      httpCookie:
//	        name: user
//	        ttl: 0s
//
// ```
type LoadBalancerSettings = v1alpha3.LoadBalancerSettings

// Consistent Hash-based load balancing can be used to provide soft
// session affinity based on HTTP headers, cookies or other
// properties. The affinity to a particular destination host may be
// lost when one or more hosts are added/removed from the destination
// service.
//
// Note: consistent hashing is less reliable at maintaining affinity than common
// "sticky sessions" implementations, which often encode a specific destination in
// a cookie, ensuring affinity is maintained as long as the backend remains.
// With consistent hash, the guarantees are weaker; any host addition or removal can
// break affinity for `1/backends` requests.
//
// Warning: consistent hashing depends on each proxy having a consistent view of endpoints.
// This is not the case when locality load balancing is enabled. Locality load balancing
// and consistent hash will only work together when all proxies are in the same locality,
// or a high level load balancer handles locality affinity.
type LoadBalancerSettings_ConsistentHashLB = v1alpha3.LoadBalancerSettings_ConsistentHashLB
type LoadBalancerSettings_ConsistentHashLB_RingHash = v1alpha3.LoadBalancerSettings_ConsistentHashLB_RingHash
type LoadBalancerSettings_ConsistentHashLB_MagLev = v1alpha3.LoadBalancerSettings_ConsistentHashLB_MagLev

// Describes a HTTP cookie that will be used as the hash key for the
// Consistent Hash load balancer.
type LoadBalancerSettings_ConsistentHashLB_HTTPCookie = v1alpha3.LoadBalancerSettings_ConsistentHashLB_HTTPCookie

// Hash based on a specific HTTP header.
type LoadBalancerSettings_ConsistentHashLB_HttpHeaderName = v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpHeaderName

// Hash based on HTTP cookie.
type LoadBalancerSettings_ConsistentHashLB_HttpCookie = v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpCookie

// Hash based on the source IP address.
// This is applicable for both TCP and HTTP connections.
type LoadBalancerSettings_ConsistentHashLB_UseSourceIp = v1alpha3.LoadBalancerSettings_ConsistentHashLB_UseSourceIp

// Hash based on a specific HTTP query parameter.
type LoadBalancerSettings_ConsistentHashLB_HttpQueryParameterName = v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpQueryParameterName

// The ring/modulo hash load balancer implements consistent hashing to backend hosts.
type LoadBalancerSettings_ConsistentHashLB_RingHash_ = v1alpha3.LoadBalancerSettings_ConsistentHashLB_RingHash_

// The Maglev load balancer implements consistent hashing to backend hosts.
type LoadBalancerSettings_ConsistentHashLB_Maglev = v1alpha3.LoadBalancerSettings_ConsistentHashLB_Maglev

// +kubebuilder:validation:XValidation:message="only one of warmupDurationSecs or warmup can be set",rule="oneof(self.warmupDurationSecs, self.warmup)"
// Standard load balancing algorithms that require no tuning.
type LoadBalancerSettings_SimpleLB = v1alpha3.LoadBalancerSettings_SimpleLB

// No load balancing algorithm has been specified by the user. Istio
// will select an appropriate default.
const LoadBalancerSettings_UNSPECIFIED LoadBalancerSettings_SimpleLB = v1alpha3.LoadBalancerSettings_UNSPECIFIED

// Deprecated. Use LEAST_REQUEST instead.
const LoadBalancerSettings_LEAST_CONN LoadBalancerSettings_SimpleLB = v1alpha3.LoadBalancerSettings_LEAST_CONN

// The random load balancer selects a random healthy host. The random
// load balancer generally performs better than round robin if no health
// checking policy is configured.
const LoadBalancerSettings_RANDOM LoadBalancerSettings_SimpleLB = v1alpha3.LoadBalancerSettings_RANDOM

// This option will forward the connection to the original IP address
// requested by the caller without doing any form of load
// balancing. This option must be used with care. It is meant for
// advanced use cases. Refer to Original Destination load balancer in
// Envoy for further details.
const LoadBalancerSettings_PASSTHROUGH LoadBalancerSettings_SimpleLB = v1alpha3.LoadBalancerSettings_PASSTHROUGH

// A basic round robin load balancing policy. This is generally unsafe
// for many scenarios (e.g. when endpoint weighting is used) as it can
// overburden endpoints. In general, prefer to use LEAST_REQUEST as a
// drop-in replacement for ROUND_ROBIN.
const LoadBalancerSettings_ROUND_ROBIN LoadBalancerSettings_SimpleLB = v1alpha3.LoadBalancerSettings_ROUND_ROBIN

// The least request load balancer spreads load across endpoints, favoring
// endpoints with the least outstanding requests. This is generally safer
// and outperforms ROUND_ROBIN in nearly all cases. Prefer to use
// LEAST_REQUEST as a drop-in replacement for ROUND_ROBIN.
const LoadBalancerSettings_LEAST_REQUEST LoadBalancerSettings_SimpleLB = v1alpha3.LoadBalancerSettings_LEAST_REQUEST

type LoadBalancerSettings_Simple = v1alpha3.LoadBalancerSettings_Simple
type LoadBalancerSettings_ConsistentHash = v1alpha3.LoadBalancerSettings_ConsistentHash
type WarmupConfiguration = v1alpha3.WarmupConfiguration

// Connection pool settings for an upstream host. The settings apply to
// each individual host in the upstream service.  See Envoy's [circuit
// breaker](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/circuit_breaking)
// for more details. Connection pool settings can be applied at the TCP
// level as well as at HTTP level.
//
// For example, the following rule sets a limit of 100 connections to redis
// service called myredissrv with a connect timeout of 30ms
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: bookinfo-redis
//
// spec:
//
//	host: myredissrv.prod.svc.cluster.local
//	trafficPolicy:
//	  connectionPool:
//	    tcp:
//	      maxConnections: 100
//	      connectTimeout: 30ms
//	      tcpKeepalive:
//	        time: 7200s
//	        interval: 75s
//
// ```
type ConnectionPoolSettings = v1alpha3.ConnectionPoolSettings

// Settings common to both HTTP and TCP upstream connections.
type ConnectionPoolSettings_TCPSettings = v1alpha3.ConnectionPoolSettings_TCPSettings

// TCP keepalive.
type ConnectionPoolSettings_TCPSettings_TcpKeepalive = v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive

// Settings applicable to HTTP1.1/HTTP2/GRPC connections.
type ConnectionPoolSettings_HTTPSettings = v1alpha3.ConnectionPoolSettings_HTTPSettings

// Policy for upgrading http1.1 connections to http2.
type ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy = v1alpha3.ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy

// Use the global default.
const ConnectionPoolSettings_HTTPSettings_DEFAULT ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy = v1alpha3.ConnectionPoolSettings_HTTPSettings_DEFAULT

// Do not upgrade the connection to http2.
// This opt-out option overrides the default.
const ConnectionPoolSettings_HTTPSettings_DO_NOT_UPGRADE ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy = v1alpha3.ConnectionPoolSettings_HTTPSettings_DO_NOT_UPGRADE

// Upgrade the connection to http2.
// This opt-in option overrides the default.
const ConnectionPoolSettings_HTTPSettings_UPGRADE ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy = v1alpha3.ConnectionPoolSettings_HTTPSettings_UPGRADE

// A Circuit breaker implementation that tracks the status of each
// individual host in the upstream service.  Applicable to both HTTP and
// TCP services.  For HTTP services, hosts that continually return 5xx
// errors for API calls are ejected from the pool for a pre-defined period
// of time. For TCP services, connection timeouts or connection
// failures to a given host counts as an error when measuring the
// consecutive errors metric. See Envoy's [outlier
// detection](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/outlier)
// for more details.
//
// The following rule sets a connection pool size of 100 HTTP1 connections
// with no more than 10 req/connection to the "reviews" service. In addition,
// it sets a limit of 1000 concurrent HTTP/2 requests and configures upstream
// hosts to be scanned every 5 mins so that any host that fails 7 consecutive
// times with a 502, 503, or 504 error code will be ejected for 15 minutes.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: reviews-cb-policy
//
// spec:
//
//	host: reviews.prod.svc.cluster.local
//	trafficPolicy:
//	  connectionPool:
//	    tcp:
//	      maxConnections: 100
//	    http:
//	      http2MaxRequests: 1000
//	      maxRequestsPerConnection: 10
//	  outlierDetection:
//	    consecutive5xxErrors: 7
//	    interval: 5m
//	    baseEjectionTime: 15m
//
// ```
type OutlierDetection = v1alpha3.OutlierDetection

// SSL/TLS related settings for upstream connections. See Envoy's [TLS
// context](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto.html#common-tls-configuration)
// for more details. These settings are common to both HTTP and TCP upstreams.
//
// For example, the following rule configures a client to use mutual TLS
// for connections to upstream database cluster.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: db-mtls
//
// spec:
//
//	host: mydbserver.prod.svc.cluster.local
//	trafficPolicy:
//	  tls:
//	    mode: MUTUAL
//	    clientCertificate: /etc/certs/myclientcert.pem
//	    privateKey: /etc/certs/client_private_key.pem
//	    caCertificates: /etc/certs/rootcacerts.pem
//
// ```
//
// The following rule configures a client to use TLS when talking to a
// foreign service whose domain matches *.foo.com.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: tls-foo
//
// spec:
//
//	host: "*.foo.com"
//	trafficPolicy:
//	  tls:
//	    mode: SIMPLE
//
// ```
//
// The following rule configures a client to use Istio mutual TLS when talking
// to rating services.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: ratings-istio-mtls
//
// spec:
//
//	host: ratings.prod.svc.cluster.local
//	trafficPolicy:
//	  tls:
//	    mode: ISTIO_MUTUAL
//
// ```
type ClientTLSSettings = v1alpha3.ClientTLSSettings

// TLS connection mode
type ClientTLSSettings_TLSmode = v1alpha3.ClientTLSSettings_TLSmode

// Do not setup a TLS connection to the upstream endpoint.
const ClientTLSSettings_DISABLE ClientTLSSettings_TLSmode = v1alpha3.ClientTLSSettings_DISABLE

// Originate a TLS connection to the upstream endpoint.
const ClientTLSSettings_SIMPLE ClientTLSSettings_TLSmode = v1alpha3.ClientTLSSettings_SIMPLE

// Secure connections to the upstream using mutual TLS by presenting
// client certificates for authentication.
const ClientTLSSettings_MUTUAL ClientTLSSettings_TLSmode = v1alpha3.ClientTLSSettings_MUTUAL

// Secure connections to the upstream using mutual TLS by presenting
// client certificates for authentication.
// Compared to Mutual mode, this mode uses certificates generated
// automatically by Istio for mTLS authentication. When this mode is
// used, all other fields in `ClientTLSSettings` should be empty.
const ClientTLSSettings_ISTIO_MUTUAL ClientTLSSettings_TLSmode = v1alpha3.ClientTLSSettings_ISTIO_MUTUAL

// Locality-weighted load balancing allows administrators to control the
// distribution of traffic to endpoints based on the localities of where the
// traffic originates and where it will terminate. These localities are
// specified using arbitrary labels that designate a hierarchy of localities in
// {region}/{zone}/{sub-zone} form. For additional detail refer to
// [Locality Weight](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/locality_weight)
// The following example shows how to setup locality weights mesh-wide.
//
// Given a mesh with workloads and their service deployed to "us-west/zone1/\*"
// and "us-west/zone2/\*". This example specifies that when traffic accessing a
// service originates from workloads in "us-west/zone1/\*", 80% of the traffic
// will be sent to endpoints in "us-west/zone1/\*", i.e the same zone, and the
// remaining 20% will go to endpoints in "us-west/zone2/\*". This setup is
// intended to favor routing traffic to endpoints in the same locality.
// A similar setting is specified for traffic originating in "us-west/zone2/\*".
//
// ```yaml
//
//	distribute:
//	  - from: us-west/zone1/*
//	    to:
//	      "us-west/zone1/*": 80
//	      "us-west/zone2/*": 20
//	  - from: us-west/zone2/*
//	    to:
//	      "us-west/zone1/*": 20
//	      "us-west/zone2/*": 80
//
// ```
//
// If the goal of the operator is not to distribute load across zones and
// regions but rather to restrict the regionality of failover to meet other
// operational requirements an operator can set a 'failover' policy instead of
// a 'distribute' policy.
//
// The following example sets up a locality failover policy for regions.
// Assume a service resides in zones within us-east, us-west & eu-west
// this example specifies that when endpoints within us-east become unhealthy
// traffic should failover to endpoints in any zone or sub-zone within eu-west
// and similarly us-west should failover to us-east.
//
// ```yaml
//
//	failover:
//	  - from: us-east
//	    to: eu-west
//	  - from: us-west
//	    to: us-east
//
// ```
type LocalityLoadBalancerSetting = v1alpha3.LocalityLoadBalancerSetting

// Describes how traffic originating in the 'from' zone or sub-zone is
// distributed over a set of 'to' zones. Syntax for specifying a zone is
// {region}/{zone}/{sub-zone} and terminal wildcards are allowed on any
// segment of the specification. Examples:
//
// `*` - matches all localities
//
// `us-west/*` - all zones and sub-zones within the us-west region
//
// `us-west/zone-1/*` - all sub-zones within us-west/zone-1
type LocalityLoadBalancerSetting_Distribute = v1alpha3.LocalityLoadBalancerSetting_Distribute

// Specify the traffic failover policy across regions. Since zone and sub-zone
// failover is supported by default this only needs to be specified for
// regions when the operator needs to constrain traffic failover so that
// the default behavior of failing over to any endpoint globally does not
// apply. This is useful when failing over traffic across regions would not
// improve service health or may need to be restricted for other reasons
// like regulatory controls.
type LocalityLoadBalancerSetting_Failover = v1alpha3.LocalityLoadBalancerSetting_Failover
