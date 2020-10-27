package edu.bridge.service.impl;

import edu.bridge.domain.Order;
import edu.bridge.mapper.OrderMapper;
import edu.bridge.service.OrderService;
import edu.bridge.service.PaymentService;
import edu.bridge.service.StorageService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
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
    public void Create(Order order, UUID uuid, int pos) {
        UUID serviceUUID = UUID.randomUUID();
        int mapperNum = 2;
        int serviceNum = 2;

        log.info("----->starting creating order<--------");
        orderMapper.create(order, uuid, pos, serviceUUID, mapperNum, serviceNum);
        log.info("----->order service beginning to call StorageService, minus count<-------");
        storageService.decrease(order.getProductId(), order.getCount(), uuid, pos);
        log.info("----->order service called StorageService, minus ended.<------");

        log.info("----->order service beginning to call AccountService, minus money.<------");
        log.info("userId:{}, money:{}", order.getUserId(), order.getMoney());
        paymentService.decrease(order.getUserId(), order.getMoney(), uuid, pos);
        log.info("----->order service called AccountService, minus ended.<------");

        log.info("----->starting modifying order<--------");
        orderMapper.update(order.getUserId(), 1, uuid, pos, serviceUUID, mapperNum, serviceNum);
        log.info("----->modifying order ended<--------");

        log.info("----->ALL HAVE BEEN DONE!<------");
    }
}
