package edu.bupt.feign;

import edu.bupt.domain.UserInfo;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;

@FeignClient(name = "user_info-service")
public interface UserInfoClient {
    @GetMapping("/createUserInfo")
    String createUserInfo(@RequestBody UserInfo userInfo);
}