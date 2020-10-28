package edu.bupt.service;

import org.springframework.web.bind.annotation.RequestParam;

import java.math.BigDecimal;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/5/2 8:02
 */
public interface PaymentService {
    void decrease(@RequestParam("userId") Long userId, @RequestParam("money")BigDecimal money);
}
