package edu.bridge.feign;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.UserInfo;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;

@FeignClient(value = "userInfo-service")
public interface UserInfoService {
    @PostMapping("/createUserInfo")
    CommonResult createUserInfo(@RequestBody UserInfo userInfo,
                                @RequestParam(required = false) CommonRequestBody commonRequestBody);
}
