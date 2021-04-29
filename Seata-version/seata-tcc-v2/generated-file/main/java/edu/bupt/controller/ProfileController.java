package edu.bupt.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.ruoyi.common.core.domain.R;
import com.ruoyi.common.core.controller.BaseController;
import edu.bupt.domain.Profile;
import edu.bupt.service.IProfileService;

/**
 * profile 提供者
 * 
 * @author bridge
 * @date 2021-04-29
 */
@RestController
@RequestMapping("profile")
public class ProfileController extends BaseController
{
	
	@Autowired
	private IProfileService profileService;
	
	/**
	 * 查询${tableComment}
	 */
	@GetMapping("get/{id}")
	public Profile get(@PathVariable("id") Long id)
	{
		return profileService.selectProfileById(id);
		
	}
	
	/**
	 * 查询profile列表
	 */
	@GetMapping("list")
	public R list(Profile profile)
	{
		startPage();
        return result(profileService.selectProfileList(profile));
	}
	
	
	/**
	 * 新增保存profile
	 */
	@PostMapping("save")
	public R addSave(@RequestBody Profile profile)
	{		
		return toAjax(profileService.insertProfile(profile));
	}

	/**
	 * 修改保存profile
	 */
	@PostMapping("update")
	public R editSave(@RequestBody Profile profile)
	{		
		return toAjax(profileService.updateProfile(profile));
	}
	
	/**
	 * 删除${tableComment}
	 */
	@PostMapping("remove")
	public R remove(String ids)
	{		
		return toAjax(profileService.deleteProfileByIds(ids));
	}
	
}
