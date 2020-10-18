package edu.bridge.controller;

import edu.bridge.domain.CommonResult;
import edu.bridge.domain.Order;
import edu.bridge.service.OrderService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.UUID;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/16 11:50
 */
@RestController
public class OrderController {

    @Resource
    private OrderService orderService;

    @GetMapping("/order/create")
    public CommonResult create(@RequestBody Order order,
                               @RequestParam(value = "UUID", required = false) UUID globalTransactionUUID,
                               @RequestParam(value = "pos", required = false) Integer pos) {
        if (globalTransactionUUID == null) {globalTransactionUUID = UUID.randomUUID();}
        if (pos == null) {pos = 0;} else {pos++;}
        orderService.Create(order, globalTransactionUUID, pos);
        return new CommonResult(200, "Order create succeeded~");
    }
}
