package edu.bridge.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.13.1)",
    comments = "Source: commonInfo.proto")
public final class CommonInfoGrpc {

  private CommonInfoGrpc() {}

  public static final String SERVICE_NAME = "CommonInfo";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<CommonInfoOuterClass.HttpRequest,
      CommonInfoOuterClass.HttpResponse> getSendToDataCenterMethod;

  public static io.grpc.MethodDescriptor<CommonInfoOuterClass.HttpRequest,
      CommonInfoOuterClass.HttpResponse> getSendToDataCenterMethod() {
    io.grpc.MethodDescriptor<CommonInfoOuterClass.HttpRequest, CommonInfoOuterClass.HttpResponse> getSendToDataCenterMethod;
    if ((getSendToDataCenterMethod = CommonInfoGrpc.getSendToDataCenterMethod) == null) {
      synchronized (CommonInfoGrpc.class) {
        if ((getSendToDataCenterMethod = CommonInfoGrpc.getSendToDataCenterMethod) == null) {
          CommonInfoGrpc.getSendToDataCenterMethod = getSendToDataCenterMethod = 
              io.grpc.MethodDescriptor.<CommonInfoOuterClass.HttpRequest, CommonInfoOuterClass.HttpResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "CommonInfo", "sendToDataCenter"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  CommonInfoOuterClass.HttpRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  CommonInfoOuterClass.HttpResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new CommonInfoMethodDescriptorSupplier("sendToDataCenter"))
                  .build();
          }
        }
     }
     return getSendToDataCenterMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static CommonInfoStub newStub(io.grpc.Channel channel) {
    return new CommonInfoStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static CommonInfoBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new CommonInfoBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static CommonInfoFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new CommonInfoFutureStub(channel);
  }

  /**
   */
  public static abstract class CommonInfoImplBase implements io.grpc.BindableService {

    /**
     */
    public void sendToDataCenter(CommonInfoOuterClass.HttpRequest request,
        io.grpc.stub.StreamObserver<CommonInfoOuterClass.HttpResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getSendToDataCenterMethod(), responseObserver);
    }

    @Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSendToDataCenterMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                CommonInfoOuterClass.HttpRequest,
                CommonInfoOuterClass.HttpResponse>(
                  this, METHODID_SEND_TO_DATA_CENTER)))
          .build();
    }
  }

  /**
   */
  public static final class CommonInfoStub extends io.grpc.stub.AbstractStub<CommonInfoStub> {
    private CommonInfoStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CommonInfoStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected CommonInfoStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CommonInfoStub(channel, callOptions);
    }

    /**
     */
    public void sendToDataCenter(CommonInfoOuterClass.HttpRequest request,
        io.grpc.stub.StreamObserver<CommonInfoOuterClass.HttpResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSendToDataCenterMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class CommonInfoBlockingStub extends io.grpc.stub.AbstractStub<CommonInfoBlockingStub> {
    private CommonInfoBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CommonInfoBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected CommonInfoBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CommonInfoBlockingStub(channel, callOptions);
    }

    /**
     */
    public CommonInfoOuterClass.HttpResponse sendToDataCenter(CommonInfoOuterClass.HttpRequest request) {
      return blockingUnaryCall(
          getChannel(), getSendToDataCenterMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class CommonInfoFutureStub extends io.grpc.stub.AbstractStub<CommonInfoFutureStub> {
    private CommonInfoFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private CommonInfoFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected CommonInfoFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new CommonInfoFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<CommonInfoOuterClass.HttpResponse> sendToDataCenter(
        CommonInfoOuterClass.HttpRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getSendToDataCenterMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SEND_TO_DATA_CENTER = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final CommonInfoImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(CommonInfoImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @Override
    @SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SEND_TO_DATA_CENTER:
          serviceImpl.sendToDataCenter((CommonInfoOuterClass.HttpRequest) request,
              (io.grpc.stub.StreamObserver<CommonInfoOuterClass.HttpResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @Override
    @SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class CommonInfoBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CommonInfoBaseDescriptorSupplier() {}

    @Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return CommonInfoOuterClass.getDescriptor();
    }

    @Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("CommonInfo");
    }
  }

  private static final class CommonInfoFileDescriptorSupplier
      extends CommonInfoBaseDescriptorSupplier {
    CommonInfoFileDescriptorSupplier() {}
  }

  private static final class CommonInfoMethodDescriptorSupplier
      extends CommonInfoBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    CommonInfoMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (CommonInfoGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new CommonInfoFileDescriptorSupplier())
              .addMethod(getSendToDataCenterMethod())
              .build();
        }
      }
    }
    return result;
  }
}
