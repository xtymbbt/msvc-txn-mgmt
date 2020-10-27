package edu.bupt.controller;

import edu.bupt.domain.CommonResult;
import edu.bupt.domain.Order;
import edu.bupt.service.OrderService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/5/1 21:57
 */
@RestController
public class OrderController {

    @Resource
    private OrderService orderService;

    @GetMapping("/order/create")
    public CommonResult create(@RequestBody Order order){
        orderService.Create(order);
        return new CommonResult(200, "Order create succeeded~");
    }
}
