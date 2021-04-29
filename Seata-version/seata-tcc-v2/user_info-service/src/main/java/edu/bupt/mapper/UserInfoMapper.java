package edu.bupt.mapper;

import edu.bupt.domain.UserInfo;
import java.util.List;

/**
 * user_infoMapper接口
 * 
 * @author bridge
 * @date 2021-04-29
 */
public interface UserInfoMapper 
{
    /**
     * 查询user_info
     * 
     * @param id user_infoID
     * @return user_info
     */
    public UserInfo selectUserInfoById(Long id);

    /**
     * 查询user_info列表
     * 
     * @param userInfo user_info
     * @return user_info集合
     */
    public List<UserInfo> selectUserInfoList(UserInfo userInfo);

    /**
     * 新增user_info
     * 
     * @param userInfo user_info
     * @return 结果
     */
    public int insertUserInfo(UserInfo userInfo);

    /**
     * 修改user_info
     * 
     * @param userInfo user_info
     * @return 结果
     */
    public int updateUserInfo(UserInfo userInfo);

    /**
     * 删除user_info
     * 
     * @param id user_infoID
     * @return 结果
     */
    public int deleteUserInfoById(Long id);

    /**
     * 批量删除user_info
     * 
     * @param ids 需要删除的数据ID
     * @return 结果
     */
    public int deleteUserInfoByIds(String[] ids);
}
