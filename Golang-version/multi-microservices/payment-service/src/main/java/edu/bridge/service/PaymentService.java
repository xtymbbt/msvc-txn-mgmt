package edu.bridge.service;

import org.springframework.web.bind.annotation.RequestParam;

import java.math.BigDecimal;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/15 11:43
 */
public interface PaymentService {
    void decrease(@RequestParam("userId") Long userId, @RequestParam("money") BigDecimal money, UUID uuid, int pos);
}
