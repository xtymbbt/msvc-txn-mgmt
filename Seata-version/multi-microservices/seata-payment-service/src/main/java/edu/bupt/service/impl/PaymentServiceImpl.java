package edu.bupt.service.impl;

import edu.bupt.dao.PaymentDao;
import edu.bupt.service.PaymentService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.math.BigDecimal;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/5/2 8:03
 */
@Service
public class PaymentServiceImpl implements PaymentService {
    private static final Logger LOGGER = LoggerFactory.getLogger(PaymentServiceImpl.class);

    @Resource
    private PaymentDao paymentDao;

    @Override
    public void decrease(Long userId, BigDecimal money) {
        LOGGER.info("------>begin minus account<-----");
        paymentDao.decrease(userId, money);
        LOGGER.info("------>minus account ended<-----");
    }
}
