package edu.bupt.feign;

import edu.bupt.domain.Profile;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;

@FeignClient(name = "profile-service")
public interface ProfileClient {
    @GetMapping("/createProfile")
    String createProfile(@RequestBody Profile profile);
}
