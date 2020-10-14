package com.bridge.controller;

import com.bridge.service.IHelloService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author Bridge Wang
 * @version 1.0
 * @date 2020/10/14 19:47
 */
@RestController
public class HelloController {
    @Autowired
    private IHelloService helloService;

    @GetMapping("/{name}")
    public String sayHello(@PathVariable String name){
        return helloService.sayHello(name);
    }
}