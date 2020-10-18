package edu.bridge.client;

import edu.bridge.grpc.CommonInfoGrpc;
import edu.bridge.grpc.CommonInfoOuterClass;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

import java.util.Map;
import java.util.UUID;
import java.util.concurrent.TimeUnit;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/15 14:44
 */
public class CommonInfoGrpcClient {
    private final ManagedChannel channel;
    private final CommonInfoGrpc.CommonInfoBlockingStub blockingStub;
    public CommonInfoGrpcClient(String host, int port){
        this(ManagedChannelBuilder.forAddress(host, port).usePlaintext());
    }

    private CommonInfoGrpcClient(ManagedChannelBuilder<?> channelBuilder){
        channel = channelBuilder.build();
        blockingStub = CommonInfoGrpc.newBlockingStub(channel);
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
    public boolean sendToDataCenter(boolean online, int pos, UUID UUID, String lastService,
                                    String currentService, String nextService,
                                    String dbName, String tableName, boolean method1,
                                    boolean method2, int query, Map<String, String> data){
        CommonInfoOuterClass.HttpRequest req = CommonInfoOuterClass.HttpRequest.newBuilder()
                .setOnline(online)
                .setPos(pos)
                .setTreeUuid(UUID.toString())
                .setLastService(lastService)
                .setCurrentService(currentService)
                .setNextService(nextService)
                .setDbName(dbName)
                .setTableName(tableName)
                .setMethod1(method1)
                .setMethod2(method2)
                .setQuery(query)
                .putAllData(data)
                .build();
        CommonInfoOuterClass.HttpResponse reply = blockingStub.sendToDataCenter(req);
        return reply.getSuccess();
    }
}
