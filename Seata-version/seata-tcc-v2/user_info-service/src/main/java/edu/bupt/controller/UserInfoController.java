package edu.bupt.controller;

import edu.bupt.domain.UserInfo;
import edu.bupt.service.UserInfoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserInfoController {
    @Autowired
    private UserInfoService userInfoService;

    @PostMapping("/createUserInfo")
    public String createUserInfo(@RequestBody UserInfo userInfo) throws Exception {
        userInfoService.create(userInfo);
        return "创建用户信息成功";
    }

}
