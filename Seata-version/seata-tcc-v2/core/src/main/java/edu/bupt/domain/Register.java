package edu.bupt.domain;


import org.apache.commons.lang.builder.ToStringBuilder;
import org.apache.commons.lang.builder.ToStringStyle;

/**
 * register对象 register
 * 
 * @author bridge
 * @date 2021-04-29
 */
public class Register
{
    private static final long serialVersionUID = 1L;

    /** primary key */
    private Long id;

    /** 用户名 */
    private String username;

    /** 密码 */
    private String password;

    /** 手机号 */
    private Long phoneNumber;

    /** 邮箱 */
    private String email;

    /** TCC事务状态 */
    private Integer status;

    public Register() {
    }

    public Register(Long id, String username, String password, Long phoneNumber, String email, Integer status) {
        this.id = id;
        this.username = username;
        this.password = password;
        this.phoneNumber = phoneNumber;
        this.email = email;
        this.status = status;
    }

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
    public void setPassword(String password) 
    {
        this.password = password;
    }

    public String getPassword() 
    {
        return password;
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
            .append("password", getPassword())
            .append("phoneNumber", getPhoneNumber())
            .append("email", getEmail())
            .append("status", getStatus())
            .toString();
    }
}
