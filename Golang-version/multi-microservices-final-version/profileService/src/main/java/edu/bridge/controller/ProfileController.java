package edu.bridge.controller;

import edu.bridge.domain.CommonRequestBody;
import edu.bridge.domain.CommonResult;
import edu.bridge.domain.Profile;
import edu.bridge.service.ProfileService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
public class ProfileController {
    @Autowired
    private ProfileService profileService;

    @PostMapping("/createProfile")
    public CommonResult createProfile(@RequestBody Profile profile,
                                      @RequestParam(required = false) CommonRequestBody commonRequestBody) {
        if (commonRequestBody == null) {
            commonRequestBody = new CommonRequestBody(UUID.randomUUID(), "root", "", "");
        }
        return profileService.createProfile(profile, commonRequestBody);
    }
}
