//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: MSAMessage.proto

package proto;

@kotlin.jvm.JvmName("-initializemSAMessage")
inline fun mSAMessage(block: proto.MSAMessageKt.Dsl.() -> kotlin.Unit): proto.MSAMessageOuterClass.MSAMessage =
  proto.MSAMessageKt.Dsl._create(proto.MSAMessageOuterClass.MSAMessage.newBuilder()).apply { block() }._build()
object MSAMessageKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  class Dsl private constructor(
    private val _builder: proto.MSAMessageOuterClass.MSAMessage.Builder
  ) {
    companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: proto.MSAMessageOuterClass.MSAMessage.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): proto.MSAMessageOuterClass.MSAMessage = _builder.build()

    /**
     * <code>.proto.MSAHeader header = 1;</code>
     */
    var header: proto.MSAHeaderOuterClass.MSAHeader
      @JvmName("getHeader")
      get() = _builder.getHeader()
      @JvmName("setHeader")
      set(value) {
        _builder.setHeader(value)
      }
    /**
     * <code>.proto.MSAHeader header = 1;</code>
     */
    fun clearHeader() {
      _builder.clearHeader()
    }
    /**
     * <code>.proto.MSAHeader header = 1;</code>
     * @return Whether the header field is set.
     */
    fun hasHeader(): kotlin.Boolean {
      return _builder.hasHeader()
    }

    /**
     * <code>string httpBody = 2;</code>
     */
    var httpBody: kotlin.String
      @JvmName("getHttpBody")
      get() = _builder.getHttpBody()
      @JvmName("setHttpBody")
      set(value) {
        _builder.setHttpBody(value)
      }
    /**
     * <code>string httpBody = 2;</code>
     */
    fun clearHttpBody() {
      _builder.clearHttpBody()
    }
  }
}
@kotlin.jvm.JvmSynthetic
inline fun proto.MSAMessageOuterClass.MSAMessage.copy(block: proto.MSAMessageKt.Dsl.() -> kotlin.Unit): proto.MSAMessageOuterClass.MSAMessage =
  proto.MSAMessageKt.Dsl._create(this.toBuilder()).apply { block() }._build()

val proto.MSAMessageOuterClass.MSAMessageOrBuilder.headerOrNull: proto.MSAHeaderOuterClass.MSAHeader?
  get() = if (hasHeader()) getHeader() else null
