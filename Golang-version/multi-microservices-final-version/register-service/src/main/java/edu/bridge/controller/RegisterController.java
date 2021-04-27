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
                                 @RequestParam(value = "globalTransactionUUID", required = false)
                                         UUID globalTransactionUUID,
                                 @RequestParam(value = "serviceUUID", required = false)
                                             String serviceUUID,
                                 @RequestParam(value = "parentUUID", required = false)
                                             String parentUUID,
                                 @RequestParam(value = "child", required = false)
                                             String child) {
        CommonRequestBody commonRequestBody;
        if (globalTransactionUUID == null) {
            commonRequestBody = new CommonRequestBody(UUID.randomUUID(), "root", "", "");
        } else {
            commonRequestBody = new CommonRequestBody(globalTransactionUUID, serviceUUID, parentUUID, child);
        }
        return registerService.registerUser(registerInfo, commonRequestBody);
    }
}
