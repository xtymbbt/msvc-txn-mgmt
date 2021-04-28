package edu.bupt.service;

import edu.bupt.tcc.ProfileTccAction;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
@Service
public class ProfileServiceImpl implements ProfileService {
    // @Autowired
    // private ProfileMapper accountMapper;

    @Autowired
    private ProfileTccAction profileTccAction;

    @Override
    public void decrease(Long userId, BigDecimal money) {
        // accountMapper.decrease(userId,money);
        profileTccAction.prepareDecreaseAccount(null, userId, money);
    }
}