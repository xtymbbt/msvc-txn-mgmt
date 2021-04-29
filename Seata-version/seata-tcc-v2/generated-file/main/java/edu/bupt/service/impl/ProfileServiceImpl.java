package edu.bupt.service.impl;

import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import edu.bupt.mapper.ProfileMapper;
import edu.bupt.domain.Profile;
import edu.bupt.service.IProfileService;
import com.ruoyi.common.core.text.Convert;

/**
 * profileService业务层处理
 * 
 * @author bridge
 * @date 2021-04-29
 */
@Service
public class ProfileServiceImpl implements IProfileService 
{
    @Autowired
    private ProfileMapper profileMapper;

    /**
     * 查询profile
     * 
     * @param id profileID
     * @return profile
     */
    @Override
    public Profile selectProfileById(Long id)
    {
        return profileMapper.selectProfileById(id);
    }

    /**
     * 查询profile列表
     * 
     * @param profile profile
     * @return profile
     */
    @Override
    public List<Profile> selectProfileList(Profile profile)
    {
        return profileMapper.selectProfileList(profile);
    }

    /**
     * 新增profile
     * 
     * @param profile profile
     * @return 结果
     */
    @Override
    public int insertProfile(Profile profile)
    {
        return profileMapper.insertProfile(profile);
    }

    /**
     * 修改profile
     * 
     * @param profile profile
     * @return 结果
     */
    @Override
    public int updateProfile(Profile profile)
    {
        return profileMapper.updateProfile(profile);
    }

    /**
     * 删除profile对象
     * 
     * @param ids 需要删除的数据ID
     * @return 结果
     */
    @Override
    public int deleteProfileByIds(String ids)
    {
        return profileMapper.deleteProfileByIds(Convert.toStrArray(ids));
    }

    /**
     * 删除profile信息
     * 
     * @param id profileID
     * @return 结果
     */
    public int deleteProfileById(Long id)
    {
        return profileMapper.deleteProfileById(id);
    }
}
