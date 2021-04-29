package edu.bupt.domain;


import org.apache.commons.lang.builder.ToStringBuilder;
import org.apache.commons.lang.builder.ToStringStyle;

/**
 * user_info对象 user_info
 * 
 * @author bridge
 * @date 2021-04-29
 */
public class UserInfo
{
    private static final long serialVersionUID = 1L;

    /** primary key */
    private Long id;

    /** 用户名 */
    private String username;

    /** 手机号 */
    private Long phoneNumber;

    /** 邮箱 */
    private String email;

    /** 照片 */
    private String photo;

    /** 经度 */
    private Long longitude;

    /** 维度 */
    private Long latitude;

    /** 当前地点 */
    private String currentAddress;

    /** 兴趣标签 */
    private String hobbyTags;

    /** 消费倾向 */
    private String consumeTags;

    /** TCC事务状态 */
    private Integer status;

    public void setId(Long id) 
    {
        this.id = id;
    }

    public Long getId() 
    {
        return id;
    }
    public void setUsername(String username) 
    {
        this.username = username;
    }

    public String getUsername() 
    {
        return username;
    }
    public void setPhoneNumber(Long phoneNumber) 
    {
        this.phoneNumber = phoneNumber;
    }

    public Long getPhoneNumber() 
    {
        return phoneNumber;
    }
    public void setEmail(String email) 
    {
        this.email = email;
    }

    public String getEmail() 
    {
        return email;
    }
    public void setPhoto(String photo) 
    {
        this.photo = photo;
    }

    public String getPhoto() 
    {
        return photo;
    }
    public void setLongitude(Long longitude) 
    {
        this.longitude = longitude;
    }

    public Long getLongitude() 
    {
        return longitude;
    }
    public void setLatitude(Long latitude) 
    {
        this.latitude = latitude;
    }

    public Long getLatitude() 
    {
        return latitude;
    }
    public void setCurrentAddress(String currentAddress) 
    {
        this.currentAddress = currentAddress;
    }

    public String getCurrentAddress() 
    {
        return currentAddress;
    }
    public void setHobbyTags(String hobbyTags) 
    {
        this.hobbyTags = hobbyTags;
    }

    public String getHobbyTags() 
    {
        return hobbyTags;
    }
    public void setConsumeTags(String consumeTags) 
    {
        this.consumeTags = consumeTags;
    }

    public String getConsumeTags() 
    {
        return consumeTags;
    }
    public void setStatus(Integer status) 
    {
        this.status = status;
    }

    public Integer getStatus() 
    {
        return status;
    }

    @Override
    public String toString() {
        return new ToStringBuilder(this, ToStringStyle.MULTI_LINE_STYLE)
            .append("id", getId())
            .append("username", getUsername())
            .append("phoneNumber", getPhoneNumber())
            .append("email", getEmail())
            .append("photo", getPhoto())
            .append("longitude", getLongitude())
            .append("latitude", getLatitude())
            .append("currentAddress", getCurrentAddress())
            .append("hobbyTags", getHobbyTags())
            .append("consumeTags", getConsumeTags())
            .append("status", getStatus())
            .toString();
    }
}
