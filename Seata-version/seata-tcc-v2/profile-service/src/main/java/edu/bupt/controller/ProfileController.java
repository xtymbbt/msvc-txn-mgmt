package edu.bupt.controller;

import edu.bupt.domain.Profile;
import edu.bupt.service.ProfileService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@Slf4j
public class ProfileController {
    @Autowired
    private ProfileService profileService;

    @PostMapping("/createProfile")
    public String createProfile(@RequestBody Profile profile) {
        profileService.createProfile(profile);
        return "创建档案成功";
    }
}
