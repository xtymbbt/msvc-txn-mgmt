package edu.bridge.service.impl;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.Order;
import edu.bridge.mapper.OrderMapper;
import edu.bridge.service.OrderService;
import edu.bridge.service.PaymentService;
import edu.bridge.service.StorageService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.HashMap;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:13
 */
@Service
@Slf4j
public class OrderServiceImpl implements OrderService {
    @Resource
    private OrderMapper orderMapper;
    @Resource
    private StorageService storageService;
    @Resource
    private PaymentService paymentService;

    @Override
    public void Create(Order order, CommonRequestBody commonRequestBody, String child) {
        HashMap<String, Boolean> children = new HashMap<>();
        if (child != null && !child.equals("")) children.put(child, true);
        children.put("storageService.decrease", true);

        log.info("----->starting creating order<--------");
        orderMapper.create(order, commonRequestBody, children);

        commonRequestBody.setParentUUID(commonRequestBody.getServiceUUID());
        commonRequestBody.setServiceUUID("storageService.decrease");
        children.remove(commonRequestBody.getServiceUUID());
        children.put("paymentService.decrease", true);
        log.info("----->order service beginning to call StorageService, minus count<-------");
        storageService.decrease(order.getProductId(), order.getCount(), commonRequestBody, "paymentService.decrease");
        log.info("----->order service called StorageService, minus ended.<------");

        commonRequestBody.setParentUUID(commonRequestBody.getServiceUUID());
        commonRequestBody.setServiceUUID("paymentService.decrease");
        children.remove(commonRequestBody.getServiceUUID());
        children.put("orderMapper.update", true);
        log.info("----->order service beginning to call AccountService, minus money.<------");
        log.info("userId:{}, money:{}", order.getUserId(), order.getMoney());
        paymentService.decrease(order.getUserId(), order.getMoney(), commonRequestBody, "orderMapper.update");
        log.info("----->order service called AccountService, minus ended.<------");

        commonRequestBody.setParentUUID(commonRequestBody.getServiceUUID());
        commonRequestBody.setServiceUUID("orderMapper.update");
        children.remove(commonRequestBody.getServiceUUID());
        log.info("----->starting modifying order<--------");
        orderMapper.update(order.getUserId(), 1, commonRequestBody, children);
        log.info("----->modifying order ended<--------");

        log.info("----->ALL HAVE BEEN DONE!<------");
    }
}
