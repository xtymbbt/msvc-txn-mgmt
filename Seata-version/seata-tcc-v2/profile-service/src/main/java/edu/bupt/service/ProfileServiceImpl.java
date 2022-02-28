package edu.bupt.service;

import edu.bupt.domain.Profile;
import edu.bupt.feign.EasyIdGeneratorClient;
import edu.bupt.tcc.ProfileTccAction;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
@Slf4j
public class ProfileServiceImpl implements ProfileService {
    @Autowired
    EasyIdGeneratorClient easyIdGeneratorClient;

    @Autowired
    private ProfileTccAction profileTccAction;

    @Override
    public void createProfile(Profile profile) {
        log.info("begin tcc prepare profile.");
        profileTccAction.prepareDecreaseAccount(null, profile);
        log.info("tcc prepare success.");
    }
}