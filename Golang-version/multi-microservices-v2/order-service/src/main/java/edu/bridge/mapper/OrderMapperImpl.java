package edu.bridge.mapper;

import edu.bridge.client.CommonInfoGrpcClient;
import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.Order;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:40
 */
@Slf4j
@Service
public class OrderMapperImpl implements OrderMapper {
    @Value("${gRPC.host}")
    private String host;
    @Value("${gRPC.port}")
    private int port;

    /**
     *     private Long id;
     *     private Long userId;
     *     private Long productId;
     *     private Integer count;
     *     private BigDecimal money;
     *     private Integer status; // order state: 1-unpaid, 2-paid
     * @param order
     */
    @Override
    public void create(Order order, CommonRequestBody commonRequestBody, HashMap<String, Boolean> children) {
        log.info("--->begin send to data center<---");
        CommonInfoGrpcClient client = new CommonInfoGrpcClient(host,port);
        Map<String, String> data = new LinkedHashMap<String, String>();
        data.put("user_id", order.getUserId().toString());
        data.put("product_id", order.getProductId().toString());
        data.put("count", order.getCount().toString());
        data.put("money", order.getMoney().toString());
        data.put("status", "0");
        /**
         * true true = 增
         * true false = 删
         * false true = 改
         * false false = 查
         */
        boolean replay = client.sendToDataCenter(true,
                commonRequestBody, children, "test_order",
                "order", true, true, "", data);
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

    @Override
    public void update(Long userId, Integer status, CommonRequestBody commonRequestBody, HashMap<String, Boolean> children) {
        log.info("--->begin send to data center<---");
        CommonInfoGrpcClient client = new CommonInfoGrpcClient(host,port);
        Map<String, String> data = new LinkedHashMap<String, String>();
        data.put("user_id", userId.toString());
        /**
         * +-*÷=共五个运算符
         * +代表数据库更新时为+
         * -代表数据库更新时为-
         * ÷代表数据库更新时为÷
         * =代表数据库更新时为赋值运算
         */
        data.put("status", "="+status);
        /**
         * true true = 增
         * true false = 删
         * false true = 改
         * false false = 查
         */
        boolean replay = client.sendToDataCenter(true,
                commonRequestBody, children, "test_order",
                "order", false, true, "user_id", data);
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
