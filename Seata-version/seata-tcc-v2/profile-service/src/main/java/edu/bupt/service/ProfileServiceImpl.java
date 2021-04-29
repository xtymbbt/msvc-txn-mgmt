package edu.bupt.service;

import edu.bupt.domain.Profile;
import edu.bupt.feign.EasyIdGeneratorClient;
import edu.bupt.tcc.ProfileTccAction;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ProfileServiceImpl implements ProfileService {
    @Autowired
    EasyIdGeneratorClient easyIdGeneratorClient;

    @Autowired
    private ProfileTccAction profileTccAction;

    @Override
    public void createProfile(Profile profile) {

        // 从全局唯一id发号器获得id
        Long id = easyIdGeneratorClient.nextId("profile_id");
        profile.setId(id);
        profileTccAction.prepareDecreaseAccount(null, profile);
    }
}