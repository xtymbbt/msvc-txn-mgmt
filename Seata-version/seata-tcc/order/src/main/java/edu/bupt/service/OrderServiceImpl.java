package edu.bupt.service;

import edu.bupt.entity.Order;
import edu.bupt.feign.PaymentClient;
import edu.bupt.feign.EasyIdGeneratorClient;
import edu.bupt.feign.StorageClient;
import edu.bupt.tcc.OrderTccAction;
import io.seata.spring.annotation.GlobalTransactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class OrderServiceImpl implements OrderService {
    // @Autowired
    // private OrderMapper orderMapper;
    @Autowired
    EasyIdGeneratorClient easyIdGeneratorClient;
    @Autowired
    private PaymentClient paymentClient;
    @Autowired
    private StorageClient storageClient;

    @Autowired
    private OrderTccAction orderTccAction;

    @GlobalTransactional
    @Override
    public void create(Order order) {
        // 从全局唯一id发号器获得id
        Long orderId = easyIdGeneratorClient.nextId("order_business");
        order.setId(orderId);

        // orderMapper.create(order);

        // 这里修改成调用 TCC 第一节端方法
        orderTccAction.prepareCreateOrder(
                null,
                order.getId(),
                order.getUserId(),
                order.getProductId(),
                order.getCount(),
                order.getMoney());


        // 修改库存
        storageClient.decrease(order.getProductId(), order.getCount());

        // 修改账户余额
        paymentClient.decrease(order.getUserId(), order.getMoney());

    }
}