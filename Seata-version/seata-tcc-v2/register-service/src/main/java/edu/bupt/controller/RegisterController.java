package edu.bupt.controller;

import edu.bupt.domain.Register;
import edu.bupt.service.RegisterService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@Slf4j
public class RegisterController {
    @Autowired
    RegisterService registerService;


    /*
    用户用这个路径进行访问：
    http://localhost:8083/register?username=zhangfe&password=123456&phoneNumber=18877925543&email=xxx@bupt.edu.cn
     */
    @GetMapping("/register")
    public String register(Register register) {
        log.info("注册用户");
        registerService.register(register);
        return "用户注册成功";
    }
}
