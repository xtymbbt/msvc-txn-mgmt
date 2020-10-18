package edu.bridge.mapper;

import edu.bridge.client.CommonInfoGrpcClient;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 13:43
 */
@Slf4j
@Service
public class StorageMapperImpl implements StorageMapper {
    @Value("${gRPC.host}")
    private String host;
    @Value("${gRPC.port}")
    private int port;

    @Override
    public void decrease(Long productId, Integer count, UUID uuid, UUID lastServiceUUID, UUID currentServiceUUID, UUID nextServiceUUID) {
        String used = "+"+count;
        String residue = "-"+count;
        log.info("--->begin send to data center<---");
        CommonInfoGrpcClient client = new CommonInfoGrpcClient(host,port);
        /**
         * true true = 增
         * true false = 删
         * false true = 改
         * false false = 查
         */
        Map<String, String> data = new LinkedHashMap<String, String>();
        data.put("product_id", productId.toString());
        data.put("used", used);
        data.put("residue", residue);
        boolean replay = client.sendToDataCenter(true, 0, uuid,
                lastServiceUUID == null ? "" : lastServiceUUID.toString(),
                currentServiceUUID == null ? "" : currentServiceUUID.toString(),
                nextServiceUUID == null ? "" : nextServiceUUID.toString(),
                "test_storage",
                "storage", false, true, 0, data);
        if (replay) {
            log.info("successfully sent to DataCenter at PORT:{}", port);
        } else {
            log.error("Data Sent to DataCenter at PORT:{} failed.", port);
        }
        try {
            client.shutdown();
        } catch (InterruptedException e) {
            log.error("channel关闭异常：err={}",e.getMessage());
        }
    }
}
