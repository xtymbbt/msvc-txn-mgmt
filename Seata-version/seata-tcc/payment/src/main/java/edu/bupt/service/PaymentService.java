package edu.bupt.service;

import java.math.BigDecimal;

public interface PaymentService {
    void decrease(Long userId, BigDecimal money);
}