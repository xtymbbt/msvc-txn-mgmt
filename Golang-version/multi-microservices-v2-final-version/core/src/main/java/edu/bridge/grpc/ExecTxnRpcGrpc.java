package edu.bridge.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.13.1)",
    comments = "Source: execTxnRpc.proto")
public final class ExecTxnRpcGrpc {

  private ExecTxnRpcGrpc() {}

  public static final String SERVICE_NAME = "ExecTxnRpc";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<ExecTxnRpcOuterClass.TxnMessage,
      ExecTxnRpcOuterClass.TxnStatus> getExecTxnMethod;

  public static io.grpc.MethodDescriptor<ExecTxnRpcOuterClass.TxnMessage,
      ExecTxnRpcOuterClass.TxnStatus> getExecTxnMethod() {
    io.grpc.MethodDescriptor<ExecTxnRpcOuterClass.TxnMessage, ExecTxnRpcOuterClass.TxnStatus> getExecTxnMethod;
    if ((getExecTxnMethod = ExecTxnRpcGrpc.getExecTxnMethod) == null) {
      synchronized (ExecTxnRpcGrpc.class) {
        if ((getExecTxnMethod = ExecTxnRpcGrpc.getExecTxnMethod) == null) {
          ExecTxnRpcGrpc.getExecTxnMethod = getExecTxnMethod = 
              io.grpc.MethodDescriptor.<ExecTxnRpcOuterClass.TxnMessage, ExecTxnRpcOuterClass.TxnStatus>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "ExecTxnRpc", "execTxn"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  ExecTxnRpcOuterClass.TxnMessage.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  ExecTxnRpcOuterClass.TxnStatus.getDefaultInstance()))
                  .setSchemaDescriptor(new ExecTxnRpcMethodDescriptorSupplier("execTxn"))
                  .build();
          }
        }
     }
     return getExecTxnMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ExecTxnRpcStub newStub(io.grpc.Channel channel) {
    return new ExecTxnRpcStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ExecTxnRpcBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new ExecTxnRpcBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ExecTxnRpcFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new ExecTxnRpcFutureStub(channel);
  }

  /**
   */
  public static abstract class ExecTxnRpcImplBase implements io.grpc.BindableService {

    /**
     */
    public void execTxn(ExecTxnRpcOuterClass.TxnMessage request,
        io.grpc.stub.StreamObserver<ExecTxnRpcOuterClass.TxnStatus> responseObserver) {
      asyncUnimplementedUnaryCall(getExecTxnMethod(), responseObserver);
    }

    @Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getExecTxnMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                ExecTxnRpcOuterClass.TxnMessage,
                ExecTxnRpcOuterClass.TxnStatus>(
                  this, METHODID_EXEC_TXN)))
          .build();
    }
  }

  /**
   */
  public static final class ExecTxnRpcStub extends io.grpc.stub.AbstractStub<ExecTxnRpcStub> {
    private ExecTxnRpcStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ExecTxnRpcStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected ExecTxnRpcStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ExecTxnRpcStub(channel, callOptions);
    }

    /**
     */
    public void execTxn(ExecTxnRpcOuterClass.TxnMessage request,
        io.grpc.stub.StreamObserver<ExecTxnRpcOuterClass.TxnStatus> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getExecTxnMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class ExecTxnRpcBlockingStub extends io.grpc.stub.AbstractStub<ExecTxnRpcBlockingStub> {
    private ExecTxnRpcBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ExecTxnRpcBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected ExecTxnRpcBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ExecTxnRpcBlockingStub(channel, callOptions);
    }

    /**
     */
    public ExecTxnRpcOuterClass.TxnStatus execTxn(ExecTxnRpcOuterClass.TxnMessage request) {
      return blockingUnaryCall(
          getChannel(), getExecTxnMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class ExecTxnRpcFutureStub extends io.grpc.stub.AbstractStub<ExecTxnRpcFutureStub> {
    private ExecTxnRpcFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private ExecTxnRpcFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @Override
    protected ExecTxnRpcFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new ExecTxnRpcFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<ExecTxnRpcOuterClass.TxnStatus> execTxn(
        ExecTxnRpcOuterClass.TxnMessage request) {
      return futureUnaryCall(
          getChannel().newCall(getExecTxnMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_EXEC_TXN = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ExecTxnRpcImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ExecTxnRpcImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @Override
    @SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_EXEC_TXN:
          serviceImpl.execTxn((ExecTxnRpcOuterClass.TxnMessage) request,
              (io.grpc.stub.StreamObserver<ExecTxnRpcOuterClass.TxnStatus>) responseObserver);
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

  private static abstract class ExecTxnRpcBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ExecTxnRpcBaseDescriptorSupplier() {}

    @Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return ExecTxnRpcOuterClass.getDescriptor();
    }

    @Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ExecTxnRpc");
    }
  }

  private static final class ExecTxnRpcFileDescriptorSupplier
      extends ExecTxnRpcBaseDescriptorSupplier {
    ExecTxnRpcFileDescriptorSupplier() {}
  }

  private static final class ExecTxnRpcMethodDescriptorSupplier
      extends ExecTxnRpcBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    ExecTxnRpcMethodDescriptorSupplier(String methodName) {
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
      synchronized (ExecTxnRpcGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ExecTxnRpcFileDescriptorSupplier())
              .addMethod(getExecTxnMethod())
              .build();
        }
      }
    }
    return result;
  }
}
