package edu.bridge.controller;

import edu.bridge.domain.CommonResult;
import edu.bridge.service.PaymentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.math.BigDecimal;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/15 11:42
 */
@RestController
public class PaymentController {
    @Autowired
    private PaymentService paymentService;

    @RequestMapping(value = "/payment/decrease")
    public CommonResult decrease(@RequestParam("userId") Long userId,
                                 @RequestParam("money") BigDecimal money,
                                 @RequestParam(value = "UUID", required = false) UUID globalTransactionUUID,
                                 @RequestParam(value = "pos", required = false) Integer pos) {
        if (globalTransactionUUID == null) {globalTransactionUUID = UUID.randomUUID();}
        if (pos == null) {pos = 0;} else {pos++;}
        paymentService.decrease(userId, money, globalTransactionUUID, pos);
        return new CommonResult(200, "success");
    }
}
