package di

import (
	"context"
	"net/url"
	"strings"
	"sync"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/i18n"
)

type contextKey string

var (
	contextKeyWidget = contextKey("widget")
)

func ContextWithWidget(ctx context.Context, widget *WidgetContainer) context.Context {
	return context.WithValue(ctx, contextKeyWidget, widget)
}

func WidgetFromContext(c context.Context) *WidgetContainer {
	if v := c.Value(contextKeyWidget); v != nil {
		return v.(*WidgetContainer)
	}

	return nil
}

type WidgetHandler interface {
	WidgetHandler(*dashboard.Response, *dashboard.Request)
	WidgetAssetFS() *assetfs.AssetFS
}

type WidgetContainerSupport interface {
	SetWidget(*WidgetContainer)
	Widget() *WidgetContainer
}

func WidgetContainerBind(bind boggart.Bind) (*WidgetContainer, bool) {
	if support, ok := bind.(WidgetContainerSupport); ok {
		container := support.Widget()
		return container, container != nil
	}

	return nil, false
}

type WidgetBind struct {
	mutex     sync.RWMutex
	container *WidgetContainer
}

func (b *WidgetBind) SetWidget(container *WidgetContainer) {
	b.mutex.Lock()
	b.container = container
	b.mutex.Unlock()
}

func (b *WidgetBind) Widget() *WidgetContainer {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	return b.container
}

type WidgetContainer struct {
	dashboard.Handler

	bind      boggart.BindItem
	configApp config.Component
}

func NewWidgetContainer(bind boggart.BindItem, configApp config.Component) *WidgetContainer {
	return &WidgetContainer{
		bind:      bind,
		configApp: configApp,
	}
}

func (c *WidgetContainer) Bind() boggart.Bind {
	return c.bind.Bind()
}

func (c *WidgetContainer) Handle(w *dashboard.Response, r *dashboard.Request) {
	switch c.bind.Status() {
	case boggart.BindStatusOnline, boggart.BindStatusOffline:
	default:
		c.NotFound(w, r)
		return
	}

	if h, ok := c.Bind().(WidgetHandler); ok {
		r = r.WithContext(ContextWithWidget(r.Context(), c))
		r = r.WithContext(boggart.ContextWithI18nDomain(r.Context(), c.I18nDomain()))
		r = r.WithContext(dashboard.ContextWithTemplateNamespace(r.Context(), c.TemplateNamespace()))
		r = r.WithContext(dashboard.ContextWithRequest(r.Context(), r))

		h.WidgetHandler(w, r)
	}
}

func (c *WidgetContainer) AssetFS() *assetfs.AssetFS {
	if w, ok := c.Bind().(WidgetHandler); ok {
		return w.WidgetAssetFS()
	}

	return nil
}

func (c *WidgetContainer) TemplateNamespace() string {
	return boggart.ComponentName + "-bind-" + c.bind.ID()
}

func (c *WidgetContainer) I18nDomain() string {
	return boggart.ComponentName + "-bind-" + c.bind.ID()
}

func (c *WidgetContainer) Translate(ctx context.Context, messageID string, context string, format ...interface{}) string {
	return i18n.Locale(ctx).Translate(boggart.I18nDomainFromContext(ctx), messageID, context, format...)
}

func (c *WidgetContainer) TranslatePlural(ctx context.Context, singleID, pluralID string, number int, context string, format ...interface{}) string {
	return i18n.Locale(ctx).TranslatePlural(boggart.I18nDomainFromContext(ctx), singleID, pluralID, number, context, format...)
}

func (c *WidgetContainer) URL(vs map[string]string) (*url.URL, error) {
	u, err := url.Parse(c.configApp.String(boggart.ConfigExternalURL))
	if err != nil {
		return nil, err
	}

	u.Path = "/" + boggart.ComponentName + "/widget/" + c.bind.ID()

	values := u.Query()

	if keysConfig := c.configApp.String(boggart.ConfigAccessKeys); keysConfig != "" {
		if keys := strings.Split(keysConfig, ","); len(keys) > 0 {
			values.Add(boggart.AccessKeyName, keys[0])
		}
	}

	for k, v := range vs {
		values.Add(k, v)
	}

	u.RawQuery = values.Encode()

	return u, nil
}

func (c *WidgetContainer) FlashError(r *dashboard.Request, messageID interface{}, context string, format ...interface{}) {
	var id string

	if e, ok := messageID.(error); ok {
		id = e.Error()
	} else {
		id = messageID.(string)
	}

	r.Session().FlashBag().Error(c.Translate(r.Context(), id, context, format...))
}

func (c *WidgetContainer) FlashSuccess(r *dashboard.Request, messageID string, context string, format ...interface{}) {
	r.Session().FlashBag().Success(c.Translate(r.Context(), messageID, context, format...))
}

func (c *WidgetContainer) FlashInfo(r *dashboard.Request, messageID string, context string, format ...interface{}) {
	r.Session().FlashBag().Info(c.Translate(r.Context(), messageID, context, format...))
}