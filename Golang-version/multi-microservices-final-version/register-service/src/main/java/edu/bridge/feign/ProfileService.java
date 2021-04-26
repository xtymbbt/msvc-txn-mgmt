package edu.bridge.feign;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.Profile;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;

@FeignClient(value = "profile-service")
public interface ProfileService {
    @PostMapping(value = "/createProfile")
    CommonResult createProfile(@RequestBody Profile profile,
                               @RequestParam(required = false) CommonRequestBody commonRequestBody);
}
