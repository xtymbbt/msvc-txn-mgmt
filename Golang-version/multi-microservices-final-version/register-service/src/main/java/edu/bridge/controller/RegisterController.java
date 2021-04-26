package edu.bridge.controller;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.RegisterInfo;
import edu.bridge.service.RegisterService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
public class RegisterController {
    @Autowired
    private RegisterService registerService;

    @GetMapping("/register")
    public CommonResult register(@RequestBody RegisterInfo registerInfo,
                                 @RequestParam(required = false) CommonRequestBody commonRequestBody) {
        if (commonRequestBody == null) {
            commonRequestBody = new CommonRequestBody(UUID.randomUUID(), "root", "", "");
        }
        return registerService.registerUser(registerInfo, commonRequestBody);
    }
}
