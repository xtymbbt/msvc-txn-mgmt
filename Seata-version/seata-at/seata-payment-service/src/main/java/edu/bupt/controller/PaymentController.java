package edu.bupt.controller;

import edu.bupt.domain.CommonResult;
import edu.bupt.service.PaymentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.math.BigDecimal;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/5/2 8:03
 */
@RestController
public class PaymentController {
    @Autowired
    private PaymentService paymentService;

    @RequestMapping(value = "/payment/decrease")
    public CommonResult decrease(@RequestParam("userId") Long userId, @RequestParam("money") BigDecimal money){
        paymentService.decrease(userId,money);
        return new CommonResult(200, "decrease payment succeeded!");
    }
}
