package edu.bupt.service;

import edu.bupt.domain.Profile;
import edu.bupt.tcc.ProfileTccAction;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ProfileServiceImpl implements ProfileService {
    // @Autowired
    // private ProfileMapper accountMapper;

    @Autowired
    private ProfileTccAction profileTccAction;

    @Override
    public void createProfile(Profile profile) {
        // accountMapper.decrease(userId,money);
        profileTccAction.prepareDecreaseAccount(null, profile);
    }
}