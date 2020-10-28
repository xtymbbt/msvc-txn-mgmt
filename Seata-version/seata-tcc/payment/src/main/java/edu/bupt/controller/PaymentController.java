package edu.bupt.controller;

import edu.bupt.service.PaymentService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.math.BigDecimal;

@RestController
@Slf4j
public class PaymentController {
    @Autowired
    private PaymentService paymentService;

    @GetMapping("/decrease")
    public String decrease(Long userId, BigDecimal money) {
        paymentService.decrease(userId,money);
        return "用户账户扣减金额成功";
    }
}
