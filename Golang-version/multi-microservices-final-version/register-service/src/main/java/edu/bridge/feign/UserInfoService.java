package edu.bridge.feign;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.UserInfo;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;

import java.util.UUID;

@FeignClient(value = "userInfo-service")
public interface UserInfoService {
    @PostMapping("/createUserInfo")
    CommonResult createUserInfo(@RequestBody UserInfo userInfo,
                                @RequestParam(value = "globalTransactionUUID", required = false)
                                        UUID globalTransactionUUID,
                                @RequestParam(value = "serviceUUID", required = false)
                                        String serviceUUID,
                                @RequestParam(value = "parentUUID", required = false)
                                        String parentUUID,
                                @RequestParam(value = "child", required = false)
                                        String child);
}
