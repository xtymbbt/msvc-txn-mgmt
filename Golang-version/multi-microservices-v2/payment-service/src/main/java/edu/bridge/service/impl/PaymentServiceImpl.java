package edu.bridge.service.impl;

import edu.bridge.client.CommonInfoGrpcClient;
import edu.bridge.domain.CommonRequestBody;
import edu.bridge.mapper.PaymentMapper;
import edu.bridge.service.PaymentService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.math.BigDecimal;
import java.util.HashMap;
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
    public void decrease(Long userId, BigDecimal money,
                         CommonRequestBody commonRequestBody,
                         String child) {
        HashMap<String, Boolean> children = new HashMap<>();
        if (child != null && !child.equals("")) children.put(child, true);
        log.info("------>begin minus account<-----");
        paymentMapper.decrease(userId, money, commonRequestBody, children);
        log.info("------>minus account ended<-----");
    }
}
