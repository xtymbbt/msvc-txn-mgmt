package edu.bridge.service;

import edu.bridge.domain.CommonResult;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;

import java.math.BigDecimal;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:02
 */
@FeignClient(value = "payment-service")
public interface PaymentService {
    @PostMapping(value = "/payment/decrease")
    CommonResult decrease(@RequestParam("userId") Long userId, @RequestParam("money") BigDecimal money,
                          @RequestParam(value = "UUID", required = false) UUID globalTransactionUUID,
                          @RequestParam(value = "pos", required = false) Integer pos);
}
