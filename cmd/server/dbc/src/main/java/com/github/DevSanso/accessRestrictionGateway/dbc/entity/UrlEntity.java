package com.github.DevSanso.accessRestrictionGateway.dbc.entity;


import javax.persistence.*;

@Entity
@Table(name="url")
public class UrlEntity {
    @Id
    public String url;
    @Column(name="access_level")
    public int accessLevel;
}
