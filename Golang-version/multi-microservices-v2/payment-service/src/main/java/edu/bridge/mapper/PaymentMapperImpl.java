package edu.bridge.mapper;

import edu.bridge.client.CommonInfoGrpcClient;
import edu.bridge.domain.CommonRequestBody;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:19
 */
@Slf4j
@Service
public class PaymentMapperImpl implements PaymentMapper {
    @Value("${gRPC.host}")
    private String host;
    @Value("${gRPC.port}")
    private int port;

    @Override
    public void decrease(Long userId, BigDecimal money, CommonRequestBody commonRequestBody, HashMap<String, Boolean> children) {
        String used = "+"+money;
        String residue = "-"+money;
        log.info("--->begin send to data center<---");
        CommonInfoGrpcClient client = new CommonInfoGrpcClient(host,port);
        /**
         * true true = 增
         * true false = 删
         * false true = 改
         * false false = 查
         */
        Map<String, String> data = new LinkedHashMap<String, String>();
        data.put("user_id", userId.toString());
        data.put("used", used);
        data.put("residue", residue);
        boolean replay = client.sendToDataCenter(true,
                commonRequestBody, children, "test_payment",
                "payment", false, true, "user_id", data);
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
