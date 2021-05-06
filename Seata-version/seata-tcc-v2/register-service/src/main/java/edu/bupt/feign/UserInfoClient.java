package edu.bupt.feign;

import edu.bupt.domain.UserInfo;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

@FeignClient(name = "user-info-service")
public interface UserInfoClient {
    @PostMapping("/createUserInfo")
    String createUserInfo(@RequestBody UserInfo userInfo);
}