package edu.bridge.service.impl;

import edu.bridge.client.CommonInfoGrpcClient;
import edu.bridge.mapper.PaymentMapper;
import edu.bridge.service.PaymentService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.math.BigDecimal;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/15 11:43
 */
@Service
@Slf4j
public class PaymentServiceImpl implements PaymentService {
    @Resource
    private PaymentMapper paymentMapper;

    @Override
    public void decrease(Long userId, BigDecimal money, UUID uuid, int pos) {
        UUID serviceUUID = UUID.randomUUID();
        log.info("------>begin minus account<-----");
        paymentMapper.decrease(userId, money, uuid, serviceUUID, 1, 0, pos);
        log.info("------>minus account ended<-----");
    }
}
