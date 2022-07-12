package com.github.DevSanso.accessRestrictionGateway.gateway.controller

import org.junit.Test
import org.junit.runner.RunWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest
import org.springframework.test.context.junit4.SpringRunner
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.status


@RunWith(SpringRunner::class)
@WebMvcTest(controllers = [RootController::class])
class RootControllerTest {
    @Autowired
    private lateinit var mvc : MockMvc


    @Test
    fun testIndexUrl() {
        mvc.perform(get("/"))
            .andExpect(status().isOk)
    }
    @Test
    fun testIndexHtmlUrl() {
        mvc.perform(get("/index.html"))
            .andExpect(status().isOk)
    }
    @Test
    fun testAuthUrl() {
        mvc.perform(get("/auth"))
            .andExpect(status().isOk)
    }
    @Test
    fun testAuthAndDeleteUrl() {
        mvc.perform(get("/auth/delete"))
            .andExpect(status().isOk)
    }
    @Test
    fun testAuthJsFileUrl() {
        mvc.perform(get("/auth/js/index.js"))
            .andExpect(status().isOk)
    }
    @Test
    fun testOutUrl() {
        mvc.perform(get("/out"))
            .andExpect(status().isOk)
    }
}