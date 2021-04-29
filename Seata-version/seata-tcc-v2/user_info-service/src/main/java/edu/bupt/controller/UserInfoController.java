package edu.bupt.controller;

import edu.bupt.domain.UserInfo;
import edu.bupt.service.UserInfoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserInfoController {
    @Autowired
    private UserInfoService userInfoService;

    @GetMapping("/createUserInfo")
    public String createUserInfo(@RequestParam UserInfo userInfo) throws Exception {
        userInfoService.create(userInfo);
        return "创建用户信息成功";
    }

}
