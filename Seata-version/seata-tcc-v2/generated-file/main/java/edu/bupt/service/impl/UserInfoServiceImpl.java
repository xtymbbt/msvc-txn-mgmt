package edu.bupt.service.impl;

import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import edu.bupt.mapper.UserInfoMapper;
import edu.bupt.domain.UserInfo;
import edu.bupt.service.IUserInfoService;
import com.ruoyi.common.core.text.Convert;

/**
 * user_infoService业务层处理
 * 
 * @author bridge
 * @date 2021-04-29
 */
@Service
public class UserInfoServiceImpl implements IUserInfoService 
{
    @Autowired
    private UserInfoMapper userInfoMapper;

    /**
     * 查询user_info
     * 
     * @param id user_infoID
     * @return user_info
     */
    @Override
    public UserInfo selectUserInfoById(Long id)
    {
        return userInfoMapper.selectUserInfoById(id);
    }

    /**
     * 查询user_info列表
     * 
     * @param userInfo user_info
     * @return user_info
     */
    @Override
    public List<UserInfo> selectUserInfoList(UserInfo userInfo)
    {
        return userInfoMapper.selectUserInfoList(userInfo);
    }

    /**
     * 新增user_info
     * 
     * @param userInfo user_info
     * @return 结果
     */
    @Override
    public int insertUserInfo(UserInfo userInfo)
    {
        return userInfoMapper.insertUserInfo(userInfo);
    }

    /**
     * 修改user_info
     * 
     * @param userInfo user_info
     * @return 结果
     */
    @Override
    public int updateUserInfo(UserInfo userInfo)
    {
        return userInfoMapper.updateUserInfo(userInfo);
    }

    /**
     * 删除user_info对象
     * 
     * @param ids 需要删除的数据ID
     * @return 结果
     */
    @Override
    public int deleteUserInfoByIds(String ids)
    {
        return userInfoMapper.deleteUserInfoByIds(Convert.toStrArray(ids));
    }

    /**
     * 删除user_info信息
     * 
     * @param id user_infoID
     * @return 结果
     */
    public int deleteUserInfoById(Long id)
    {
        return userInfoMapper.deleteUserInfoById(id);
    }
}
