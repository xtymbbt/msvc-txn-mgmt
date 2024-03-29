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
 * <pre>
 *表明我们定义了一个命名为Simple的服务(接口)，内部有一个远程rpc方法，名字为SayHello。
 *我们只要在server端实现这个接口，在实现类中书写我们的业务代码。在client端调用这个接口。
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.13.1)",
    comments = "Source: simple.proto")
public final class SimpleGrpc {

  private SimpleGrpc() {}

  public static final String SERVICE_NAME = "Simple";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<SimpleOuterClass.HelloRequest,
      SimpleOuterClass.HelloReplay> getSayHelloMethod;

  public static io.grpc.MethodDescriptor<SimpleOuterClass.HelloRequest,
      SimpleOuterClass.HelloReplay> getSayHelloMethod() {
    io.grpc.MethodDescriptor<SimpleOuterClass.HelloRequest, SimpleOuterClass.HelloReplay> getSayHelloMethod;
    if ((getSayHelloMethod = SimpleGrpc.getSayHelloMethod) == null) {
      synchronized (SimpleGrpc.class) {
        if ((getSayHelloMethod = SimpleGrpc.getSayHelloMethod) == null) {
          SimpleGrpc.getSayHelloMethod = getSayHelloMethod = 
              io.grpc.MethodDescriptor.<SimpleOuterClass.HelloRequest, SimpleOuterClass.HelloReplay>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "Simple", "SayHello"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  SimpleOuterClass.HelloRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  SimpleOuterClass.HelloReplay.getDefaultInstance()))
                  .setSchemaDescriptor(new SimpleMethodDescriptorSupplier("SayHello"))
                  .build();
          }
        }
     }
     return getSayHelloMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SimpleStub newStub(io.grpc.Channel channel) {
    return new SimpleStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SimpleBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new SimpleBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SimpleFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new SimpleFutureStub(channel);
  }

  /**
   * <pre>
   *表明我们定义了一个命名为Simple的服务(接口)，内部有一个远程rpc方法，名字为SayHello。
   *我们只要在server端实现这个接口，在实现类中书写我们的业务代码。在client端调用这个接口。
   * </pre>
   */
  public static abstract class SimpleImplBase implements io.grpc.BindableService {

    /**
     */
    public void sayHello(SimpleOuterClass.HelloRequest request,
        io.grpc.stub.StreamObserver<SimpleOuterClass.HelloReplay> responseObserver) {
      asyncUnimplementedUnaryCall(getSayHelloMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSayHelloMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                SimpleOuterClass.HelloRequest,
                SimpleOuterClass.HelloReplay>(
                  this, METHODID_SAY_HELLO)))
          .build();
    }
  }

  /**
   * <pre>
   *表明我们定义了一个命名为Simple的服务(接口)，内部有一个远程rpc方法，名字为SayHello。
   *我们只要在server端实现这个接口，在实现类中书写我们的业务代码。在client端调用这个接口。
   * </pre>
   */
  public static final class SimpleStub extends io.grpc.stub.AbstractStub<SimpleStub> {
    private SimpleStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SimpleStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SimpleStub(channel, callOptions);
    }

    /**
     */
    public void sayHello(SimpleOuterClass.HelloRequest request,
        io.grpc.stub.StreamObserver<SimpleOuterClass.HelloReplay> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSayHelloMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   *表明我们定义了一个命名为Simple的服务(接口)，内部有一个远程rpc方法，名字为SayHello。
   *我们只要在server端实现这个接口，在实现类中书写我们的业务代码。在client端调用这个接口。
   * </pre>
   */
  public static final class SimpleBlockingStub extends io.grpc.stub.AbstractStub<SimpleBlockingStub> {
    private SimpleBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SimpleBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SimpleBlockingStub(channel, callOptions);
    }

    /**
     */
    public SimpleOuterClass.HelloReplay sayHello(SimpleOuterClass.HelloRequest request) {
      return blockingUnaryCall(
          getChannel(), getSayHelloMethod(), getCallOptions(), request);
    }
  }

  /**
   * <pre>
   *表明我们定义了一个命名为Simple的服务(接口)，内部有一个远程rpc方法，名字为SayHello。
   *我们只要在server端实现这个接口，在实现类中书写我们的业务代码。在client端调用这个接口。
   * </pre>
   */
  public static final class SimpleFutureStub extends io.grpc.stub.AbstractStub<SimpleFutureStub> {
    private SimpleFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private SimpleFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SimpleFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new SimpleFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<SimpleOuterClass.HelloReplay> sayHello(
        SimpleOuterClass.HelloRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getSayHelloMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SAY_HELLO = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final SimpleImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(SimpleImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SAY_HELLO:
          serviceImpl.sayHello((SimpleOuterClass.HelloRequest) request,
              (io.grpc.stub.StreamObserver<SimpleOuterClass.HelloReplay>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class SimpleBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    SimpleBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return SimpleOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Simple");
    }
  }

  private static final class SimpleFileDescriptorSupplier
      extends SimpleBaseDescriptorSupplier {
    SimpleFileDescriptorSupplier() {}
  }

  private static final class SimpleMethodDescriptorSupplier
      extends SimpleBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    SimpleMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (SimpleGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new SimpleFileDescriptorSupplier())
              .addMethod(getSayHelloMethod())
              .build();
        }
      }
    }
    return result;
  }
}
