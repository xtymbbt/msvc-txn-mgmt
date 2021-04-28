package edu.bupt.service;

import edu.bupt.tcc.UserInfoTccAction;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserInfoServiceImpl implements UserInfoService {
    // @Autowired
    // private UserInfoMapper storageMapper;

    @Autowired
    private UserInfoTccAction userInfoTccAction;

    @Override
    public void decrease(Long productId, Integer count) throws Exception {
        // storageMapper.decrease(productId,count);
        userInfoTccAction.prepareDecreaseStorage(null, productId, count);
    }

}
