package edu.bupt.controller;

import edu.bupt.service.UserInfoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserInfoController {
    @Autowired
    private UserInfoService userInfoService;

    @GetMapping("/decrease")
    public String decrease(Long productId, Integer count) throws Exception {
        userInfoService.decrease(productId,count);
        return "减少商品库存成功";
    }

}
