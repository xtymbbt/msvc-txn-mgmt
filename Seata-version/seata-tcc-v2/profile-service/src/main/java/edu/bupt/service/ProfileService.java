package edu.bupt.service;

import java.math.BigDecimal;

public interface ProfileService {
    void decrease(Long userId, BigDecimal money);
}