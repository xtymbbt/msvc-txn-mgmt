package edu.bridge.client;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.grpc.ExecTxnRpcGrpc;
import edu.bridge.grpc.ExecTxnRpcOuterClass;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/15 14:44
 */
@Component
public class ExecTxnRpcGrpcClient {
    @Value("${gRPC.host}")
    private String host;
    @Value("${gRPC.port}")
    private int port;

    private ManagedChannel channel;
    private ExecTxnRpcGrpc.ExecTxnRpcBlockingStub blockingStub;

    @PostConstruct
    public void init() {
        ManagedChannelBuilder<?> channelBuilder =
                ManagedChannelBuilder.forAddress(host, port).usePlaintext();
        channel = channelBuilder.build();
        blockingStub = ExecTxnRpcGrpc.newBlockingStub(channel);
    }

    public void shutdown()throws InterruptedException{
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }

    /**
     *   bool online = 1;
     *   uint32 pos = 2;
     *   bool is_the_last_service = 3;
     *   string db_name = 4;
     *   string table_name = 5;
     *   bool method1 = 6;
     *   bool method2 = 7;
     *   map<string, string> data = 8;
     */
    public boolean sendToDataCenter(boolean online, CommonRequestBody commonRequestBody,
                                    List<String> children,
                                    String dbName, String tableName, boolean method1,
                                    boolean method2, String query, Map<String, String> data){
        ExecTxnRpcOuterClass.TxnMessage req = ExecTxnRpcOuterClass.TxnMessage.newBuilder()
                .setOnline(online)
                .setTreeUuid(commonRequestBody.getGlobalTransactionUUID().toString())
                .setServiceUuid(commonRequestBody.getServiceUUID())
                .setParentUuid(commonRequestBody.getParentUUID())
                .addAllChildren(children)
                .setDbName(dbName)
                .setTableName(tableName)
                .setMethod1(method1)
                .setMethod2(method2)
                .setQuery(query)
                .putAllData(data)
                .build();
        ExecTxnRpcOuterClass.TxnStatus reply = blockingStub.execTxn(req);
        return reply.getStatus() == 200;
    }
}
