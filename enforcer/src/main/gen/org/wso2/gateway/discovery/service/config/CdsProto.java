// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/service/config/cds.proto

package org.wso2.gateway.discovery.service.config;

public final class CdsProto {
  private CdsProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\'wso2/discovery/service/config/cds.prot" +
      "o\022\030discovery.service.config\032*envoy/servi" +
      "ce/discovery/v3/discovery.proto2\373\001\n\026Conf" +
      "igDiscoveryService\022r\n\rStreamConfigs\022,.en" +
      "voy.service.discovery.v3.DiscoveryReques" +
      "t\032-.envoy.service.discovery.v3.Discovery" +
      "Response\"\000(\0010\001\022m\n\014FetchConfigs\022,.envoy.s" +
      "ervice.discovery.v3.DiscoveryRequest\032-.e" +
      "nvoy.service.discovery.v3.DiscoveryRespo" +
      "nse\"\000B\200\001\n)org.wso2.gateway.discovery.ser" +
      "vice.configB\010CdsProtoP\001ZDgithub.com/envo" +
      "yproxy/go-control-plane/wso2/discovery/s" +
      "ervice/config\210\001\001b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          io.envoyproxy.envoy.service.discovery.v3.DiscoveryProto.getDescriptor(),
        });
    io.envoyproxy.envoy.service.discovery.v3.DiscoveryProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
