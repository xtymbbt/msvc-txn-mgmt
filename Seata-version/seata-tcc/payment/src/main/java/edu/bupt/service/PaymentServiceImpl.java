package edu.bupt.service;

import edu.bupt.tcc.PaymentTccAction;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
@Service
public class PaymentServiceImpl implements PaymentService {
    // @Autowired
    // private PaymentMapper accountMapper;

    @Autowired
    private PaymentTccAction paymentTccAction;

    @Override
    public void decrease(Long userId, BigDecimal money) {
        // accountMapper.decrease(userId,money);
        paymentTccAction.prepareDecreaseAccount(null, userId, money);
    }
}