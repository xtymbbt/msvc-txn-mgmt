package edu.bridge.controller;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.UserInfo;
import edu.bridge.service.UserInfoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
public class UserInfoController {
    @Autowired
    private UserInfoService userInfoService;

    @PostMapping("/createUserInfo")
    public CommonResult createUserInfo(@RequestBody UserInfo userInfo,
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
        return userInfoService.recordUserInfo(userInfo, commonRequestBody);
    }
}
